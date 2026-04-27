package http

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/ricardovalehrio/dio-expert-session-finance/adapter/http/actuator"
	transaction "github.com/ricardovalehrio/dio-expert-session-finance/adapter/http/transactions"
)

// Init blablabla
func Init() {
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateATransactions)

	http.HandleFunc("/health", actuator.Health)

	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":8080", nil)
}
