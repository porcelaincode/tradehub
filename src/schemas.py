from pydantic import BaseModel
from typing import Optional, List, Dict

class InstrumentSubscription(BaseModel):
    instrument_token: str
    exchange_segment: str

class OrderCreate(BaseModel):
    amo: Optional[str] = "NO"
    disclosed_quantity: Optional[str] = "0"
    exchange_segment: str
    market_protection: Optional[str] = "0"
    product: str
    pf: Optional[str] = "N"
    price: Optional[str]
    order_type: str
    quantity: str
    validity: str
    trigger_price: Optional[str]
    trading_symbol: str
    transaction_type: str
    tag: Optional[str]

class OrderResponse(BaseModel):
    type: str
    data: Dict

class PositionData(BaseModel):
    buyAmt: str
    cfSellAmt: str
    prod: str
    exSeg: str
    sqrFlg: str
    actId: str
    cfBuyQty: str
    cfSellQty: str
    tok: str
    flBuyQty: str
    flSellQty: str
    sellAmt: str
    posFlg: str
    cfBuyAmt: str
    stkPrc: str
    trdSym: str
    sym: str
    expDt: str
    type: str
    series: str
    brdLtQty: str
    exp: str
    optTp: str
    genNum: str
    genDen: str
    prcNum: str
    prcDen: str
    lotSz: str
    multiplier: str
    precision: str
    hsUpTm: str

class PositionResponse(BaseModel):
    stat: str
    stCode: int
    data: List[PositionData]

