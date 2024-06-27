package models

type Order struct {
    AM                string  `json:"am,omitempty"`
    DisclosedQuantity int     `json:"disclosed_quantity,omitempty"`
    ExchangeSegment   string  `json:"exchange_segment"`
    MarketProtection  int     `json:"market_protection,omitempty"`
    Product           string  `json:"product"`
    PF                string  `json:"pf,omitempty"`
    Price             float64 `json:"price,omitempty"`
    OrderType         string  `json:"order_type"`
    Quantity          int     `json:"quantity"`
    Validity          string  `json:"validity"`
    TriggerPrice      float64 `json:"trigger_price,omitempty"`
    TradingSymbol     string  `json:"trading_symbol"`
    TransactionType   string  `json:"transaction_type"`
    Tag               string  `json:"tag,omitempty"`
}
