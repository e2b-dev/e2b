import logging
from typing import Optional, Callable, Any, Tuple, List, Union

from pydantic import BaseModel, PrivateAttr

from e2b import EnvVars, SyncSession

logger = logging.getLogger(__name__)


class Artifact(BaseModel):
    name: str
    _session: SyncSession = PrivateAttr()

    def __init__(self, **data: Any):
        super().__init__(**data)
        self._session = data["_session"]

    def __hash__(self):
        return hash(self.name)

    def read(self) -> bytes:
        return self._session.download_file(self.name)

    def download(self, path: Optional[str] = None) -> None:
        data = self.read()
        with open(path or self.name, "wb") as f:
            f.write(data)


class DataAnalysis(SyncSession):
    env_id = "Python3-DataAnalysis"

    def __init__(
        self,
        api_key: Optional[str] = None,
        cwd: Optional[str] = None,
        env_vars: Optional[EnvVars] = None,
        on_stdout: Optional[Callable[[str], Any]] = None,
        on_stderr: Optional[Callable[[str], Any]] = None,
        on_artifact: Optional[Callable[[Artifact], Any]] = None,
        on_exit: Optional[Callable[[int], Any]] = None,
    ):
        self.on_artifact = on_artifact
        super().__init__(
            id=self.env_id,
            api_key=api_key,
            cwd=cwd,
            env_vars=env_vars,
            on_stdout=on_stdout,
            on_stderr=on_stderr,
            on_exit=on_exit,
        )
        self.open()

    def create(self, *args, **kwargs):
        raise Exception("Wrong syntax. Use only `DataAnalysis(...)`")

    def run_python(
        self,
        code: str,
        on_stdout: Optional[Callable[[str], Any]] = None,
        on_stderr: Optional[Callable[[str], Any]] = None,
        on_exit: Optional[Callable[[int], Any]] = None,
    ) -> Tuple[str, str, List[Artifact]]:
        artifacts = set()

        def register_artifacts(event: Any):
            if event.operation == "Create":
                artifact = Artifact(name=event.path, _session=self)
                artifacts.add(artifact)
                if self.on_artifact:
                    try:
                        self.on_artifact(artifact)
                    except Exception as e:
                        logger.error("Failed to process artifact", exc_info=e)

        watcher = self.filesystem.watch_dir("/home/user/artifacts")
        watcher.add_event_listener(register_artifacts)
        watcher.start()

        process = self.process.start(
            f'python -c "{code}"',
            on_stdout=on_stdout,
            on_stderr=on_stderr,
            on_exit=on_exit,
        )
        process.wait()

        watcher.stop()

        return process.output.stdout, process.output.stderr, list(artifacts)

    def install_python_package(self, package_names: Union[str, List[str]]):
        if isinstance(package_names, list):
            package_names = " ".join(package_names)

        process = self.process.start(f"pip install {package_names}")
        process.wait()

        if process.exit_code != 0:
            raise Exception(
                f"Failed to install package {package_names}: {process.output.stderr}"
            )

    def install_system_package(self, package_names: Union[str, List[str]]):
        if isinstance(package_names, list):
            package_names = " ".join(package_names)

        process = self.process.start(f"apt-get {package_names}")
        process.wait()

        if process.exit_code != 0:
            raise Exception(
                f"Failed to install package {package_names}: {process.output.stderr}"
            )