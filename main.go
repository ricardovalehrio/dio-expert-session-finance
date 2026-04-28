package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/ricardovalehrio/dio-expert-session-finance/model/transaction"
)

func main() {
	http.HandleFunc("/transactions", getTransactions)
	http.HandleFunc("/transactions/create", createATransaction)

	fmt.Println("Servidor rodando na porta 8080...")
	http.ListenAndServe(":8080", nil)
}

func getTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	layout := "2006-01-02T15:04:05"
	salaryReceived, _ := time.Parse(layout, "2026-04-26T14:56:05")

	minhasTransacoes := transaction.Transactions{
		{
			Title:     "Salário Mensal",
			Amount:    1200.0,
			Type:      0,
			CreatedAt: salaryReceived,
		},
	}
	json.NewEncoder(w).Encode(minhasTransacoes)
}

func createATransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var novaTransacao transaction.Transaction
	body, err := ioutil.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		w.WriteHeader(http.StatusBadRequest) // Erro 400 se o body estiver vazio
		fmt.Fprint(w, "Corpo da requisição vazio")
		return
	}

	err = json.Unmarshal(body, &novaTransacao)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "JSON inválido")
		return
	}

	fmt.Printf("Recebido: %+v\n", novaTransacao)
	w.WriteHeader(http.StatusCreated)
}
