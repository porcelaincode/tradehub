from sqlalchemy import Column, Integer, String, Float, JSON, Date
from database import Base

class Instrument(Base):
    __tablename__ = 'instruments'
    instrument_token = Column(Integer, primary_key=True, index=True)
    exchange_token = Column(Integer)
    tradingsymbol = Column(String)
    name = Column(String)
    last_price = Column(Float)
    expiry = Column(Date)
    strike = Column(Float)
    tick_size = Column(Float)
    lot_size = Column(Integer)
    instrument_type = Column(String)
    segment = Column(String)
    exchange = Column(String)

class Order(Base):
    __tablename__ = "orders"
    id = Column(Integer, primary_key=True, index=True)
    amo = Column(String, default="NO")
    disclosed_quantity = Column(String, default="0")
    exchange_segment = Column(String)
    market_protection = Column(String, default="0")
    product = Column(String)
    pf = Column(String, default="N")
    price = Column(String)
    order_type = Column(String)
    quantity = Column(String)
    validity = Column(String)
    trigger_price = Column(String)
    trading_symbol = Column(String)
    transaction_type = Column(String)
    tag = Column(String, default=None)

class Position(Base):
    __tablename__ = "positions"
    id = Column(Integer, primary_key=True, index=True)
    instrument_token = Column(Integer)
    buyAmt = Column(String)
    cfSellAmt = Column(String)
    prod = Column(String)
    exSeg = Column(String)
    sqrFlg = Column(String)
    actId = Column(String)
    cfBuyQty = Column(String)
    cfSellQty = Column(String)
    tok = Column(String)
    flBuyQty = Column(String)
    flSellQty = Column(String)
    sellAmt = Column(String)
    posFlg = Column(String)
    cfBuyAmt = Column(String)
    stkPrc = Column(String)
    trdSym = Column(String)
    sym = Column(String)
    expDt = Column(String)
    type = Column(String)
    series = Column(String)
    brdLtQty = Column(String)
    exp = Column(String)
    optTp = Column(String)
    genNum = Column(String)
    genDen = Column(String)
    prcNum = Column(String)
    prcDen = Column(String)
    lotSz = Column(String)
    multiplier = Column(String)
    precision = Column(String)
    hsUpTm = Column(String)

