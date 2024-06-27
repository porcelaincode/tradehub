package models

type Position struct {
    ID          int    `json:"id"`
    BuyAmt      string `json:"buyAmt"`
    CfSellAmt   string `json:"cfSellAmt"`
    Prod        string `json:"prod"`
    ExSeg       string `json:"exSeg"`
    SqrFlg      string `json:"sqrFlg"`
    ActId       string `json:"actId"`
    CfBuyQty    string `json:"cfBuyQty"`
    CfSellQty   string `json:"cfSellQty"`
    Tok         string `json:"tok"`
    FlBuyQty    string `json:"flBuyQty"`
    FlSellQty   string `json:"flSellQty"`
    SellAmt     string `json:"sellAmt"`
    PosFlg      string `json:"posFlg"`
    CfBuyAmt    string `json:"cfBuyAmt"`
    StkPrc      string `json:"stkPrc"`
    TrdSym      string `json:"trdSym"`
    Sym         string `json:"sym"`
    ExpDt       string `json:"expDt"`
    Type        string `json:"type"`
    Series      string `json:"series"`
    BrdLtQty    string `json:"brdLtQty"`
    Exp         string `json:"exp"`
    OptTp       string `json:"optTp"`
    GenNum      string `json:"genNum"`
    GenDen      string `json:"genDen"`
    PrcNum      string `json:"prcNum"`
    PrcDen      string `json:"prcDen"`
    LotSz       string `json:"lotSz"`
    Multiplier  string `json:"multiplier"`
    Precision   string `json:"precision"`
    HsUpTm      string `json:"hsUpTm"`
}