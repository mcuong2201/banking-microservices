package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"risk-detection/internal/kafka"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8084"
	}

	go kafka.StartConsumer()

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK")
	})

	log.Printf("Risk Detection service starting on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
