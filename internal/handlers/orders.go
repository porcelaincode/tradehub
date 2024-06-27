package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/porcelaincode/tradehub/internal/database"
	"github.com/porcelaincode/tradehub/internal/models"
	wsManager "github.com/porcelaincode/tradehub/internal/websocket"
)

func PlaceOrderHandler(w http.ResponseWriter, r *http.Request) {
    var order models.Order
    err := json.NewDecoder(r.Body).Decode(&order)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    err = PlaceOrder(order)
    if err != nil {
        http.Error(w, "Failed to place order", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

func PlaceOrder(order models.Order) error {
    _, err := database.DB.Exec(`
        INSERT INTO orders (
            am, disclosed_quantity, exchange_segment, market_protection, product, pf, price, order_type,
            quantity, validity, trigger_price, trading_symbol, transaction_type, tag
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
        order.AM, order.DisclosedQuantity, order.ExchangeSegment, order.MarketProtection, order.Product, order.PF,
        order.Price, order.OrderType, order.Quantity, order.Validity, order.TriggerPrice, order.TradingSymbol,
        order.TransactionType, order.Tag)

    if err != nil {
        log.Printf("Error placing order: %v", err)
        return err
    }

    orderEvent := map[string]interface{}{
        "type": "order",
        "data": order,
    }
    message, err := json.Marshal(orderEvent)
    if err != nil {
        log.Printf("Error marshalling order event: %v", err)
        return err
    }

    wsManager.GetManager().Broadcast(message)
    return nil
}