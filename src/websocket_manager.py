from fastapi import WebSocket
from typing import List
import asyncio

class WebSocketManager:
    def __init__(self):
        self.active_connections: List[WebSocket] = []
        self.subscribed_tokens: List[str] = []

    async def connect(self, websocket: WebSocket):
        await websocket.accept()
        self.active_connections.append(websocket)

    async def disconnect(self, websocket: WebSocket):
        self.active_connections.remove(websocket)

    async def disconnect_all(self):
        for connection in self.active_connections:
            await connection.close()
        self.active_connections.clear()

    async def broadcast(self, message: str):
        for connection in self.active_connections:
            await connection.send_text(message)

    def subscribe(self, token: str):
        if token not in self.subscribed_tokens:
            self.subscribed_tokens.append(token)

    def unsubscribe(self, token: str):
        if token in self.subscribed_tokens:
            self.subscribed_tokens.remove(token)

    def emit_order_event(self, order):
        event = {
            "type": "order",
            "data": order.__dict__
        }
        asyncio.create_task(self.broadcast(str(event)))

    def emit_position_event(self, position):
        event = {
            "type": "position",
            "data": position.__dict__
        }
        asyncio.create_task(self.broadcast(str(event)))

websocket_manager = WebSocketManager()

