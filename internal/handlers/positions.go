package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/porcelaincode/tradehub/internal/services"
)

func GetPositionsHandler(w http.ResponseWriter, r *http.Request) {
    positions, err := services.GetPositions()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    response := map[string]interface{}{
        "stat": "Ok",
        "stCode": 200,
        "data": positions,
    }
    json.NewEncoder(w).Encode(response)
}