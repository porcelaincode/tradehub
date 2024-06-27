package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/porcelaincode/tradehub/internal/database"
	"github.com/porcelaincode/tradehub/internal/handlers"
	"github.com/porcelaincode/tradehub/internal/services"

	wsManager "github.com/porcelaincode/tradehub/internal/websocket"
)

func main() {
    database.InitDB("tradehub.db")

    if err := services.DownloadAndSaveInstruments(); err != nil {
        log.Fatalf("Error downloading instruments: %v", err)
    }

    r := mux.NewRouter()
    r.HandleFunc("/orders", handlers.PlaceOrderHandler).Methods("POST")
    r.HandleFunc("/positions", handlers.GetPositionsHandler).Methods("GET")
    r.HandleFunc("/subscribe", handlers.SubscribeHandler).Methods("POST")
    r.HandleFunc("/ws", handlers.WebSocketHandler)

    go wsManager.GetManager().Start()

    log.Println("Server started on :8090")
    log.Fatal(http.ListenAndServe(":8090", r))
}