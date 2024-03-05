import binascii
import json
import os
import uuid
from time import sleep
from typing import Optional, Callable, Any, List

import requests
from pydantic import BaseModel
from websocket import create_connection

from e2b import Sandbox, EnvVars, ProcessMessage
from e2b.constants import TIMEOUT


class Result(BaseModel):
    output: Optional[str] = None
    stdout: List[str] = []
    stderr: List[str] = []
    error: Optional[str] = None
    # TODO: This will be changed in the future, it's just to enable the use of display_data
    display_data: List[dict] = []


class CodeInterpreterV2(Sandbox):
    template = "code-interpreter-stateful"

    def __init__(
        self,
        api_key: Optional[str] = None,
        cwd: Optional[str] = None,
        env_vars: Optional[EnvVars] = None,
        timeout: Optional[float] = TIMEOUT,
        on_stdout: Optional[Callable[[ProcessMessage], Any]] = None,
        on_stderr: Optional[Callable[[ProcessMessage], Any]] = None,
        on_exit: Optional[Callable[[int], Any]] = None,
        **kwargs,
    ):
        super().__init__(
            template=self.template,
            api_key=api_key,
            cwd=cwd,
            env_vars=env_vars,
            timeout=timeout,
            on_stdout=on_stdout,
            on_stderr=on_stderr,
            on_exit=on_exit,
            **kwargs,
        )

        self._jupyter_server_token = binascii.hexlify(os.urandom(24)).decode("ascii")
        self._jupyter_kernel_id = self._start_jupyter()

    def _start_jupyter(self) -> str:
        self.process.start(
            f"jupyter server --IdentityProvider.token={self._jupyter_server_token}"
        )

        url = f"{self.get_protocol()}://{self.get_hostname(8888)}"
        headers = {"Authorization": f"Token {self._jupyter_server_token}"}

        response = requests.get(url, headers=headers)
        while response.status_code != 200:
            sleep(0.2)
            response = requests.get(f"{url}/api", headers=headers)

        response = requests.post(f"{url}/api/kernels", headers=headers)
        if response.status_code != 201:
            raise Exception(f"Error creating kernel: {response.status_code}")

        kernel = json.loads(response.text)
        return kernel["id"]

    def _connect_kernel(self):
        header = {"Authorization": f"Token {self._jupyter_server_token}"}
        return create_connection(
            f"{self.get_protocol('ws')}://{self.get_hostname(8888)}/api/kernels/{self._jupyter_kernel_id}/channels",
            header=header,
        )

    @staticmethod
    def _send_execute_request(code: str):
        msg_id = str(uuid.uuid4())
        session = str(uuid.uuid4())

        return {
            "header": {
                "msg_id": msg_id,
                "username": "e2b",
                "session": session,
                "msg_type": "execute_request",
                "version": "5.3",
            },
            "parent_header": {},
            "metadata": {},
            "content": {
                "code": code,
                "silent": False,
                "store_history": False,
                "user_expressions": {},
                "allow_stdin": False,
            },
        }

    @staticmethod
    def _wait_for_result(ws) -> Result:
        result = Result()
        was_busy = False

        while True:
            response = json.loads(ws.recv())
            if response["msg_type"] == "error":
                result.error = response["content"]["traceback"]
            elif response["msg_type"] == "stream":
                if response["content"]["name"] == "stdout":
                    result.stdout.append(response["content"]["text"])
                elif response["content"]["name"] == "stderr":
                    result.stderr.append(response["content"]["text"])

            elif response["msg_type"] == "display_data":
                result.display_data.append(response["content"]["data"])
            elif response["msg_type"] == "execute_result":
                result.output = response["content"]["data"]["text/plain"]

            elif response["msg_type"] == "status":
                if response["content"]["execution_state"] == "idle":
                    if was_busy:
                        break
                elif response["content"]["execution_state"] == "error":
                    result.error = "An error occurred while executing the code"
                    break
                elif response["content"]["execution_state"] == "busy":
                    if not was_busy:
                        was_busy = True

        return result

    def execute(self, code: str) -> Result:
        ws = self._connect_kernel()
        ws.send(json.dumps(self._send_execute_request(code)))
        result = self._wait_for_result(ws)

        ws.close()

        return result
