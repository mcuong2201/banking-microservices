package main

import (
	"log"
	"net/http"
	"transaction-ledger/internal/handler"
	"transaction-ledger/internal/kafka"
)

func main() {
	kafka.InitKafkaWriter("localhost:9092", "transactions")

	http.HandleFunc("/transaction", handler.TransactionHandler)

	log.Println("🚀 transaction-ledger running on :8080")
	log.Fatal(http.ListenAndServe(":8083", nil))
}
