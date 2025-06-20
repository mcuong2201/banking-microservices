package risk

import (
	"sync"
	"time"
)

type Transaction struct {
	TransactionID string    `json:"transaction_id"`
	FromUserID    string    `json:"from_user_id"`
	ToUserID      string    `json:"to_user_id"`
	Amount        int64     `json:"amount"`
	Timestamp     time.Time `json:"timestamp"`
}

type RiskEvaluator struct {
	mu            sync.Mutex
	userTxHistory map[string][]time.Time
}

func NewRiskEvaluator() *RiskEvaluator {
	return &RiskEvaluator{
		userTxHistory: make(map[string][]time.Time),
	}
}

func (r *RiskEvaluator) Evaluate(tx Transaction) string {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	history := r.userTxHistory[tx.FromUserID]

	var recentTx []time.Time
	for _, t := range history {
		if now.Sub(t) < time.Hour {
			recentTx = append(recentTx, t)
		}
	}

	recentTx = append(recentTx, now)
	r.userTxHistory[tx.FromUserID] = recentTx

	if len(recentTx) > 5 || tx.Amount > 500_000_000 {
		return "HIGH_RISK"
	}

	return "LOW_RISK"
}
