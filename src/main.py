import asyncio
from fastapi import FastAPI, WebSocket, Depends, HTTPException
from sqlalchemy.orm import Session
import uvicorn
from database import engine, SessionLocal, Base
from models import Instrument
from schemas import InstrumentSubscription, OrderCreate, OrderResponse, PositionResponse
from crud import create_order, get_positions, update_position, download_instruments
from websocket_manager import websocket_manager
from mock_data_generator import start_mock_data_generator

app = FastAPI()

Base.metadata.create_all(bind=engine)

@app.on_event("startup")
async def startup_event():
    await download_instruments()
    asyncio.create_task(start_mock_data_generator())

@app.on_event("shutdown")
async def shutdown_event():
    await websocket_manager.disconnect_all()

@app.post("/subscribe")
async def subscribe_instruments(subscriptions: list[InstrumentSubscription]):
    for subscription in subscriptions:
        websocket_manager.subscribe(subscription.instrument_token)
    return {"message": "Subscribed successfully"}

@app.post("/orders", response_model=OrderResponse)
async def create_new_order(order: OrderCreate, db: Session = Depends(SessionLocal)):
    try:
        return create_order(db, order)
    except Exception as e:
        raise HTTPException(status_code=400, detail=str(e))

@app.get("/positions", response_model=PositionResponse)
async def get_all_positions(db: Session = Depends(SessionLocal)):
    return get_positions(db)

@app.websocket("/ws")
async def websocket_endpoint(websocket: WebSocket):
    await websocket_manager.connect(websocket)
    try:
        while True:
            await websocket.receive_text()
    except Exception as e:
        print(f"Connection error: {e}")
    finally:
        await websocket_manager.disconnect(websocket)

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8000)

