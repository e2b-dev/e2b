from io import TextIOBase
from typing import Iterator, List, Literal, Optional, overload
from e2b.sandbox.filesystem.filesystem import WriteData, WriteEntry

import e2b_connect
import httpcore
import httpx
from e2b.connection_config import (
    ConnectionConfig,
    Username,
    KEEPALIVE_PING_HEADER,
    KEEPALIVE_PING_INTERVAL_SEC,
)
from e2b.envd.api import ENVD_API_FILES_ROUTE, handle_envd_api_exception
from e2b.envd.filesystem import filesystem_connect, filesystem_pb2
from e2b.envd.rpc import authentication_header, handle_rpc_exception
from e2b.exceptions import SandboxException
from e2b.sandbox.filesystem.filesystem import EntryInfo, map_file_type
from e2b.sandbox_sync.filesystem.watch_handle import WatchHandle


class Filesystem:
    def __init__(
        self,
        envd_api_url: str,
        connection_config: ConnectionConfig,
        pool: httpcore.ConnectionPool,
        envd_api: httpx.Client,
    ) -> None:
        self._envd_api_url = envd_api_url
        self._connection_config = connection_config
        self._pool = pool
        self._envd_api = envd_api

        self._rpc = filesystem_connect.FilesystemClient(
            envd_api_url,
            # TODO: Fix and enable compression again — the headers compression is not solved for streaming.
            # compressor=e2b_connect.GzipCompressor,
            pool=pool,
            json=True,
        )

    @overload
    def read(
        self,
        path: str,
        format: Literal["text"] = "text",
        user: Username = "user",
        request_timeout: Optional[float] = None,
    ) -> str: ...

    @overload
    def read(
        self,
        path: str,
        format: Literal["bytes"],
        user: Username = "user",
        request_timeout: Optional[float] = None,
    ) -> bytearray: ...

    @overload
    def read(
        self,
        path: str,
        format: Literal["stream"],
        user: Username = "user",
        request_timeout: Optional[float] = None,
    ) -> Iterator[bytes]: ...

    def read(
        self,
        path: str,
        format: Literal["text", "bytes", "stream"] = "text",
        user: Username = "user",
        request_timeout: Optional[float] = None,
    ):
        """Read from file"""
        r = self._envd_api.get(
            ENVD_API_FILES_ROUTE,
            params={"path": path, "username": user},
            timeout=self._connection_config.get_request_timeout(request_timeout),
        )

        err = handle_envd_api_exception(r)
        if err:
            raise err

        if format == "text":
            return r.text
        elif format == "bytes":
            return bytearray(r.content)
        elif format == "stream":
            return r.iter_bytes()

    @overload
    def write(
        self,
        path: str,
        data: WriteData,
        user: Username = "user",
        request_timeout: Optional[float] = None,
    ) -> EntryInfo:
        """Write to file"""

    @overload
    def write(
        self,
        files: List[WriteEntry],
        user: Optional[Username] = "user",
        path: Optional[str] = None,
        request_timeout: Optional[float] = None,
    ) -> List[EntryInfo]:
        """Write multiple files"""

    def write(
        self,
        path_or_files: str | List[WriteEntry],
        data_or_user: WriteData | Username = "user",
        user_or_request_timeout: Optional[float | Username] = None,
        request_timeout_or_none: Optional[float] = None
    ) -> EntryInfo | List[EntryInfo]:
        """Write to file(s)
        When writing to a file that doesn't exist, the file will get created.
        When writing to a file that already exists, the file will get overwritten.
        When writing to a file that's in a directory that doesn't exist, you'll get an error.
        """

        path, write_files, user, request_timeout,  = None, [], "user", None
        if isinstance(path_or_files, str):
            if isinstance(data_or_user, list):
                raise Exception("Cannot specify path with array of files")
            path, write_files, user, request_timeout = \
                path_or_files, [{"path": path_or_files, "data": data_or_user}], user_or_request_timeout or "user", request_timeout_or_none
        else:
            path, write_files, user, request_timeout = \
                None, path_or_files, data_or_user, user_or_request_timeout
        
        if len(write_files) == 0:
            raise Exception("Need at least one file to write")

        # Prepare the files for the multipart/form-data request
        httpx_files = []
        for file in write_files:
            file_path, file_data = file['path'], file['data']
            if isinstance(file_data, str) or isinstance(file_data, bytes):
                httpx_files.append(('file', (file_path, file_data)))
            elif isinstance(file_data, TextIOBase):
                httpx_files.append(('file', (file_path, file_data.read())))
            else:
                raise ValueError(f"Unsupported data type for file {file_path}")

        params = {"username": user}
        if path is not None: params["path"] = path

        r = self._envd_api.post(
            ENVD_API_FILES_ROUTE,
            files=httpx_files,
            params=params,
            timeout=self._connection_config.get_request_timeout(request_timeout),
        )

        err = handle_envd_api_exception(r)
        if err:
            raise err

        write_files = r.json()

        if not isinstance(write_files, list) or len(write_files) == 0:
            raise Exception("Expected to receive information about written file")

        if len(write_files) == 1 and path:
            file = write_files[0]
            return EntryInfo(**file)
        else:
            return [EntryInfo(**file) for file in write_files]

    def list(
        self,
        path: str,
        user: Username = "user",
        request_timeout: Optional[float] = None,
    ) -> List[EntryInfo]:
        """List directory"""
        try:
            res = self._rpc.list_dir(
                filesystem_pb2.ListDirRequest(path=path),
                request_timeout=self._connection_config.get_request_timeout(
                    request_timeout
                ),
                headers=authentication_header(user),
            )

            entries: List[EntryInfo] = []
            for entry in res.entries:
                event_type = map_file_type(entry.type)

                if event_type:
                    entries.append(
                        EntryInfo(name=entry.name, type=event_type, path=entry.path)
                    )

            return entries
        except Exception as e:
            raise handle_rpc_exception(e)

    def exists(
        self,
        path: str,
        user: Username = "user",
        request_timeout: Optional[float] = None,
    ) -> bool:
        """Check if file exists."""
        try:
            self._rpc.stat(
                filesystem_pb2.StatRequest(path=path),
                request_timeout=self._connection_config.get_request_timeout(
                    request_timeout
                ),
                headers=authentication_header(user),
            )
            return True

        except Exception as e:
            if isinstance(e, e2b_connect.ConnectException):
                if e.status == e2b_connect.Code.not_found:
                    return False
            raise handle_rpc_exception(e)

    def remove(
        self,
        path: str,
        user: Username = "user",
        request_timeout: Optional[float] = None,
    ) -> None:
        """Remove file"""
        try:
            self._rpc.remove(
                filesystem_pb2.RemoveRequest(path=path),
                request_timeout=self._connection_config.get_request_timeout(
                    request_timeout
                ),
                headers=authentication_header(user),
            )
        except Exception as e:
            raise handle_rpc_exception(e)

    def rename(
        self,
        old_path: str,
        new_path: str,
        user: Username = "user",
        request_timeout: Optional[float] = None,
    ) -> EntryInfo:
        """Rename file"""
        try:
            r = self._rpc.move(
                filesystem_pb2.MoveRequest(
                    source=old_path,
                    destination=new_path,
                ),
                request_timeout=self._connection_config.get_request_timeout(
                    request_timeout
                ),
                headers=authentication_header(user),
            )

            return EntryInfo(
                name=r.entry.name,
                type=map_file_type(r.entry.type),
                path=r.entry.path,
            )
        except Exception as e:
            raise handle_rpc_exception(e)

    def make_dir(
        self,
        path: str,
        user: Username = "user",
        request_timeout: Optional[float] = None,
    ) -> bool:
        """Create directory and all parent directories"""
        try:
            self._rpc.make_dir(
                filesystem_pb2.MakeDirRequest(path=path),
                request_timeout=self._connection_config.get_request_timeout(
                    request_timeout
                ),
                headers=authentication_header(user),
            )

            return True
        except Exception as e:
            if isinstance(e, e2b_connect.ConnectException):
                if e.status == e2b_connect.Code.already_exists:
                    return False
            raise handle_rpc_exception(e)

    def watch(
        self,
        path: str,
        user: Username = "user",
        request_timeout: Optional[float] = None,
        timeout: Optional[float] = 60,
    ):
        """Watch directory for changes."""
        events = self._rpc.watch_dir(
            filesystem_pb2.WatchDirRequest(path=path),
            request_timeout=self._connection_config.get_request_timeout(
                request_timeout
            ),
            timeout=timeout,
            headers={
                **authentication_header(user),
                KEEPALIVE_PING_HEADER: str(KEEPALIVE_PING_INTERVAL_SEC),
            },
        )

        try:
            start_event = events.__next__()

            if not start_event.HasField("start"):
                raise SandboxException(
                    f"Failed to start watch: expected start event, got {start_event}",
                )

            return WatchHandle(events=events)
        except Exception as e:
            raise handle_rpc_exception(e)
