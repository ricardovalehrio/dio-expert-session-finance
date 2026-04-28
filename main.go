package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	// IMPORTANTE: Este caminho deve ser igual ao do seu go.mod
	"github.com/ricardovalehrio/dio-expert-session-finance/model/transaction"
)

func main() {
	// Rotas da API
	http.HandleFunc("/transactions", getTransactions)
	http.HandleFunc("/transactions/create", createATransaction)

	fmt.Println("Servidor iniciado na porta 8080...")
	// Inicia o servidor e escuta na porta 8080
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

	// CORREÇÃO: Note o uso de 'transaction.' antes dos nomes das structs
	var transactions = transaction.Transactions{
		transaction.Transaction{
			Title:     "Salário",
			Amount:    1200.0,
			Type:      0,
			CreatedAt: salaryReceived,
		},
	}

	_ = json.NewEncoder(w).Encode(transactions)
}

func createATransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// CORREÇÃO: Usando o prefixo do pacote aqui também
	var res = transaction.Transactions{}
	body, _ := ioutil.ReadAll(r.Body)

	_ = json.Unmarshal(body, &res)

	fmt.Printf("Recebido via Postman: %+v\n", res)
}
