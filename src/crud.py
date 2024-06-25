import requests
from sqlalchemy.orm import Session
from models import Instrument, Order, Position
from schemas import OrderCreate, OrderResponse, PositionResponse
from websocket_manager import websocket_manager
from database import SessionLocal
import csv

def download_instruments():
    url = 'http://api.kite.trade/instruments'
    response = requests.get(url)
    response.raise_for_status()
    
    decoded_content = response.content.decode('utf-8')
    csv_reader = csv.DictReader(decoded_content.splitlines(), delimiter=',')

    db: Session = SessionLocal()
    try:
        for row in csv_reader:
            instrument = Instrument(
                instrument_token=int(row['instrument_token']),
                exchange_token=int(row['exchange_token']),
                tradingsymbol=row['tradingsymbol'],
                name=row['name'],
                last_price=float(row['last_price']),
                expiry=row['expiry'],
                strike=float(row['strike']),
                tick_size=float(row['tick_size']),
                lot_size=int(row['lot_size']),
                instrument_type=row['instrument_type'],
                segment=row['segment'],
                exchange=row['exchange']
            )
            db.add(instrument)
        db.commit()
    finally:
        db.close()

def create_order(db: Session, order: OrderCreate):
    db_order = Order(**order.dict())
    db.add(db_order)
    db.commit()
    db.refresh(db_order)
    update_position(db, order)
    websocket_manager.emit_order_event(db_order)
    return OrderResponse(type="order", data=db_order.__dict__)

def get_positions(db: Session):
    positions = db.query(Position).all()
    return PositionResponse(stat="Ok", stCode=200, data=positions)

def update_position(db: Session, order: OrderCreate):
    position_data = {
        "instrument_token": order.trading_symbol,
        "buyAmt": "1000.00",  # Simulate position data
        "cfSellAmt": "0.00",
        "prod": order.product,
        "exSeg": order.exchange_segment,
        "sqrFlg": "Y",
        "actId": "PRS2206",
        "cfBuyQty": "0",
        "cfSellQty": "0",
        "tok": order.trading_symbol,
        "flBuyQty": "25",
        "flSellQty": "25",
        "sellAmt": "1000.00",
        "posFlg": "true",
        "cfBuyAmt": "0.00",
        "stkPrc": "0.00",
        "trdSym": order.trading_symbol,
        "sym": order.trading_symbol,
        "expDt": "29 Jul, 2021",
        "type": "FUTIDX",
        "series": "XX",
        "brdLtQty": "25",
        "exp": "1627569000",
        "optTp": "XX",
        "genNum": "1",
        "genDen": "1",
        "prcNum": "1",
        "prcDen": "1",
        "lotSz": "25",
        "multiplier": "1",
        "precision": "2",
        "hsUpTm": "2021/07/13 18:34:44"
    }

    position = db.query(Position).filter(Position.instrument_token == order.trading_symbol).first()
    if position:
        for key, value in position_data.items():
            setattr(position, key, value)
    else:
        position = Position(**position_data)
        db.add(position)
    db.commit()
    websocket_manager.emit_position_event(position)

