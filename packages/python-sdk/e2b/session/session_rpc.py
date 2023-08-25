import json
import logging
from concurrent.futures import ThreadPoolExecutor
from queue import Queue
from threading import Event
from typing import Any, Callable, Dict, Iterator, List

from e2b.session.exception import RpcException
from e2b.session.websocket_client import Message, Notification, start_websocket
from e2b.utils.future import DeferredFuture, run_async_func_in_new_loop
from jsonrpcclient import Error, Ok, request_json
from jsonrpcclient.id_generators import decimal as decimal_id_generator
from pydantic import BaseModel, PrivateAttr

logger = logging.getLogger(__name__)


import asyncio


class Event_ts(asyncio.Event):
    # TODO: clear() method
    def set(self):
        # FIXME: The _loop attribute is not documented as public api!
        self._loop.call_soon_threadsafe(super().set)


def to_response_or_notification(response: Dict[str, Any]) -> Message:
    """Create a Response namedtuple from a dict"""
    logger.info(f"Received response: {response}")
    if "error" in response:
        return Error(
            response["error"]["code"],
            response["error"]["message"],
            response["error"].get("data"),
            response["id"],
        )
    elif "result" in response and "id" in response:
        return Ok(response["result"], response["id"])

    elif "params" in response:
        return Notification(method=response["method"], params=response["params"])

    raise ValueError("Invalid response", response)


class SessionRpc(BaseModel):
    url: str
    on_message: Callable[[Notification], None]

    _id_generator: Iterator[int] = PrivateAttr(default_factory=decimal_id_generator)
    _waiting_for_replies: Dict[int, DeferredFuture] = PrivateAttr(default_factory=dict)
    _queue_in: Queue = PrivateAttr(default_factory=Queue)
    _queue_out: Queue = PrivateAttr(default_factory=Queue)
    _process_cleanup: List[Callable[[], Any]] = PrivateAttr(default_factory=list)

    class Config:
        arbitrary_types_allowed = True

    async def process_messages(self):
        while True:
            data = await self._queue_out.get()
            message = to_response_or_notification(json.loads(data))

            logger.info(f"Current waiting handlers: {self._waiting_for_replies}")
            if isinstance(message, Ok):
                if (
                    message.id in self._waiting_for_replies
                    and self._waiting_for_replies[message.id]
                ):
                    self._waiting_for_replies[message.id](message.result)
                    return
            elif isinstance(message, Error):
                if (
                    message.id in self._waiting_for_replies
                    and self._waiting_for_replies[message.id]
                ):
                    self._waiting_for_replies[message.id].reject(
                        RpcException(
                            code=message.code,
                            message=message.message,
                            id=message.id,
                            data=message.data,
                        )
                    )
                    return

            elif isinstance(message, Notification):
                self.on_message(message)

    async def connect(self):
        started = Event_ts()
        cancelled = Event_ts()
        task = asyncio.create_task(self.process_messages())
        executor = ThreadPoolExecutor(max_workers=1, thread_name_prefix="e2b_ws")
        websocket_task = asyncio.get_running_loop().run_in_executor(
            executor,
            run_async_func_in_new_loop,
            start_websocket(
                self.url,
                self.on_message,
                self._queue_in,
                self._queue_out,
                self._waiting_for_replies,
                started,
                cancelled,
            ),
        )
        self._process_cleanup.append(task.cancel)
        self._process_cleanup.append(cancelled.set)
        self._process_cleanup.append(websocket_task.cancel)
        self._process_cleanup.append(executor.shutdown)
        await started.wait()

    async def send_message(self, method: str, params: List[Any]) -> Any:
        id = next(self._id_generator)
        request = request_json(method, params, id)
        future_reply = DeferredFuture(self._process_cleanup)

        try:
            self._waiting_for_replies[id] = future_reply
            logger.info(f"Queueing: {request}")
            self._queue_in.put(request)
            logger.info(f"Queue size: {self._queue_in.qsize()}")
            logger.info(f"Waiting for reply: {request}")
            r = await future_reply
            return r
        except Exception as e:
            logger.error(f"Error: {request} {e}")
            raise e
        finally:
            del self._waiting_for_replies[id]
            logger.info(f"Removed waiting handler for {id}")

    def _close(self):
        for cancel in self._process_cleanup:
            cancel()

        self._process_cleanup.clear()

        for handler in self._waiting_for_replies.values():
            handler.cancel()
            del handler

    async def close(self):
        self._close()
