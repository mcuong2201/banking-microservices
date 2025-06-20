package handler

import (
    "encoding/json"
    "net/http"
    "strconv"
    "risk-detection/internal/app"
)

func RiskHandler(w http.ResponseWriter, r *http.Request) {
    amountStr := r.URL.Query().Get("amount")
    amount, err := strconv.ParseFloat(amountStr, 64)
    if err != nil {
        http.Error(w, "Invalid amount", http.StatusBadRequest)
        return
    }

    risk := app.EvaluateRisk(amount)
    json.NewEncoder(w).Encode(map[string]string{
        "risk": risk,
    })
}
