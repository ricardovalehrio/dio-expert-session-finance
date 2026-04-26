package http

import (
	"net/http"

	"github.com/ricardovalehrio/dio-expert-session-finance/adapter/http/actuator"
	transaction "github.com/ricardovalehrio/dio-expert-session-finance/adapter/http/transactions"
)

// Init blablabla
func Init() {
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateATransactions)

	http.HandleFunc("/health", actuator.Health)

	http.ListenAndServe(":8080", nil)
}
