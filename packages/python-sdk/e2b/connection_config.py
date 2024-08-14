import os

from typing import Literal, Optional

REQUEST_TIMEOUT: float = 30.0  # 30s


class ConnectionConfig:
    @staticmethod
    def _domain():
        return os.getenv("E2B_DOMAIN", "e2b.dev")

    @staticmethod
    def _debug():
        return os.getenv("E2B_DEBUG", "false") == "true"

    @staticmethod
    def _api_key():
        return os.getenv("E2B_API_KEY")

    @staticmethod
    def _access_token():
        return os.getenv("E2B_ACCESS_TOKEN")

    def __init__(
        self,
        domain: Optional[str] = None,
        debug: Optional[bool] = None,
        api_key: Optional[str] = None,
        access_token: Optional[str] = None,
        request_timeout: Optional[float] = None,
    ):
        self.domain = domain or ConnectionConfig._domain()
        self.debug = debug or ConnectionConfig._debug()
        self.api_key = api_key or ConnectionConfig._api_key()
        self.access_token = access_token or ConnectionConfig._access_token()

        self.request_timeout = ConnectionConfig._get_request_timeout(
            REQUEST_TIMEOUT,
            request_timeout,
        )

        if request_timeout == 0:
            self.request_timeout = None
        elif request_timeout is not None:
            self.request_timeout = request_timeout
        else:
            self.request_timeout = REQUEST_TIMEOUT

        self.api_url = (
            "http://localhost:3000" if self.debug else f"https://api.{self.domain}"
        )

    @staticmethod
    def _get_request_timeout(
        default_timeout: Optional[float],
        request_timeout: Optional[float],
    ):
        if request_timeout == 0:
            return None
        elif request_timeout is not None:
            return request_timeout
        else:
            return default_timeout

    def get_request_timeout(self, request_timeout: Optional[float] = None):
        return self._get_request_timeout(self.request_timeout, request_timeout)


Username = Literal["root", "user"]

default_username: Username = "user"