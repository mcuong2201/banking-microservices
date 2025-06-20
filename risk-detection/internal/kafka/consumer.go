package kafka

import (
	"context"
	"encoding/json"
	"log"
	"risk-detection/internal/risk"
	"time"

	"github.com/segmentio/kafka-go"
)

var evaluator *risk.RiskEvaluator

func StartConsumer() {
	evaluator = risk.NewRiskEvaluator()

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "transactions",
		GroupID: "risk-detection-group",
	})

	log.Println("Kafka consumer started...")

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Kafka read error: %v", err)
			time.Sleep(time.Second)
			continue
		}

		var tx risk.Transaction
		if err := json.Unmarshal(m.Value, &tx); err != nil {
			log.Printf("Failed to unmarshal transaction: %v", err)
			continue
		}

		riskLevel := evaluator.Evaluate(tx)
		log.Printf("🚨 Risk level: %s - User: %s - Amount: %.2f - Time: %s",
			riskLevel, tx.FromUserID, float64(tx.Amount)/100, tx.Timestamp.Format(time.RFC3339))
	}
}
