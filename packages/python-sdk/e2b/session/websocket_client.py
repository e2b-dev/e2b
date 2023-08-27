import asyncio
import logging
from queue import Queue
from threading import Event
from typing import Any, Callable, List

from janus import SyncQueue as JanusQueue
from websockets import WebSocketClientProtocol, connect
from websockets.exceptions import ConnectionClosed
from websockets.typing import Data

logger = logging.getLogger(__name__)


class WebSocket:
    def __init__(
        self,
        url: str,
        started: Event,
        stopped: Event,
        queue_in: Queue[dict],
        queue_out: JanusQueue[Data],
    ):
        self._ws: WebSocketClientProtocol | None = None
        self.url = url
        self.started = started
        self.stopped = stopped
        self._process_cleanup: List[Callable[[], Any]] = []
        self._queue_in = queue_in
        self._queue_out = queue_out

    async def run(self):
        await self._connect()
        await self.close()

    async def _send_message(self):
        logger.info("Starting to send messages")
        while True:
            if self._queue_in.empty():
                await asyncio.sleep(0)
                continue
            message = self._queue_in.get()
            logger.debug(f"Got message: {message}")
            if self._ws:
                await self._ws.send(message)
                self._queue_in.task_done()
            else:
                logger.error("No websocket connection")

    async def _receive_message(self):
        try:
            async for message in self._ws:
                logger.debug(f"Received message: {message}")
                self._queue_out.put(message)

        except Exception as e:
            logger.error(f"Error: {e}")

    async def _connect(self):
        try:
            async for websocket in connect(self.url, max_size=None, max_queue=None):
                try:
                    self._ws = websocket
                    self.started.set()
                    logger.info(f"Connected to {self.url}")

                    send_task = asyncio.create_task(
                        self._send_message(), name="send_message"
                    )
                    self._process_cleanup.append(send_task.cancel)

                    receive_task = asyncio.create_task(
                        self._receive_message(), name="receive_message"
                    )
                    self._process_cleanup.append(receive_task.cancel)

                    while not self.stopped.is_set():
                        await asyncio.sleep(0)

                    logger.info("WebSocket stopped")
                    break
                except ConnectionClosed:
                    logger.warning("Reconnecting...")
                    continue
        except BaseException as e:
            logger.error(f"Error: {e}")

    def _close(self):
        for cancel in self._process_cleanup:
            cancel()

        self._process_cleanup.clear()

    async def close(self):
        self._close()

        if self._ws:
            await self._ws.close()

    @classmethod
    async def start(
        cls,
        url,
        queue_in: Queue[dict],
        queue_out: JanusQueue[Data],
        started: Event,
        stopped: Event,
    ):
        websocket = cls(
            url=url,
            stopped=stopped,
            started=started,
            queue_in=queue_in,
            queue_out=queue_out,
        )
        await websocket.run()
