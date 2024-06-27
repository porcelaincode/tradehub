package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/porcelaincode/tradehub/internal/models"
)

var subscribedTokens = make(map[int]bool)

func SubscribeHandler(w http.ResponseWriter, r *http.Request) {
    var subscriptions []models.Subscription
    if err := json.NewDecoder(r.Body).Decode(&subscriptions); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

    for _, subscription := range subscriptions {
        subscribedTokens[subscription.InstrumentToken] = true
    }

    w.WriteHeader(http.StatusOK)
}
