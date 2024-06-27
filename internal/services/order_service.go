package services

import (
	"log"

	"github.com/porcelaincode/tradehub/internal/database"
	"github.com/porcelaincode/tradehub/internal/models"
)

func PlaceOrder(order models.Order) error {
    _, err := database.DB.Exec(`
        INSERT INTO orders (
            am, disclosed_quantity, exchange_segment, market_protection, product, pf, price, order_type, quantity, validity, trigger_price, trading_symbol, transaction_type, tag
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
        order.AM, order.DisclosedQuantity, order.ExchangeSegment, order.MarketProtection, order.Product, order.PF, order.Price, order.OrderType, order.Quantity, order.Validity,
        order.TriggerPrice, order.TradingSymbol, order.TransactionType, order.Tag)
    if err != nil {
        log.Printf("Error placing order: %v", err)
        return err
    }

    // (WebSocket event logic will be implemented later)
		
    return nil
}

func GetPositions() ([]models.Position, error) {
    rows, err := database.DB.Query(`SELECT * FROM positions`)
    if err != nil {
        log.Printf("Error retrieving positions: %v", err)
        return nil, err
    }
    defer rows.Close()

    var positions []models.Position
    for rows.Next() {
        var position models.Position
        err := rows.Scan(&position.ID, &position.BuyAmt, &position.CfSellAmt, &position.Prod, &position.ExSeg, &position.SqrFlg, &position.ActId, &position.CfBuyQty, &position.CfSellQty,
            &position.Tok, &position.FlBuyQty, &position.FlSellQty, &position.SellAmt, &position.PosFlg, &position.CfBuyAmt, &position.StkPrc, &position.TrdSym, &position.Sym, &position.ExpDt,
            &position.Type, &position.Series, &position.BrdLtQty, &position.Exp, &position.OptTp, &position.GenNum, &position.GenDen, &position.PrcNum, &position.PrcDen, &position.LotSz,
            &position.Multiplier, &position.Precision, &position.HsUpTm)
        if err != nil {
            log.Printf("Error scanning position: %v", err)
            continue
        }
        positions = append(positions, position)
    }
    return positions, nil
}