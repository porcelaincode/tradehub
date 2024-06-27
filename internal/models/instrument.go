package models

import (
	"time"
)

type Instrument struct {
    InstrumentToken int       `json:"instrument_token"`
    ExchangeToken   int       `json:"exchange_token"`
    TradingSymbol   string    `json:"tradingsymbol"`
    Name            string    `json:"name"`
    LastPrice       float64   `json:"last_price"`
    Expiry          time.Time `json:"expiry"`
    Strike          float64   `json:"strike"`
    TickSize        float64   `json:"tick_size"`
    LotSize         int       `json:"lot_size"`
    InstrumentType  string    `json:"instrument_type"`
    Segment         string    `json:"segment"`
    Exchange        string    `json:"exchange"`
}