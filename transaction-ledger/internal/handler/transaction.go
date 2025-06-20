package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"transaction-ledger/internal/kafka"

	"github.com/google/uuid"
)

type Transaction struct {
	TransactionID string    `json:"transaction_id"`
	FromUserID    string    `json:"from_user_id"`
	ToUserID      string    `json:"to_user_id"`
	Amount        int64     `json:"amount"`
	Timestamp     time.Time `json:"timestamp"`
}

func TransactionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var tx Transaction
	if err := json.NewDecoder(r.Body).Decode(&tx); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	tx.TransactionID = uuid.New().String()
	tx.Timestamp = time.Now().UTC()

	go kafka.PublishMessage(tx.TransactionID, tx)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"transaction_id": tx.TransactionID,
	})
}
