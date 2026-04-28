// Arquivo: main.go (na raiz do projeto)
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	// O caminho deve bater exatamente com o module no seu go.mod
	"github.com/ricardovalehrio/dio-expert-session-finance/model/transaction"
)

func main() {
	// Definindo as rotas
	http.HandleFunc("/transactions", getTransactions)
	http.HandleFunc("/transactions/create", createATransaction)

	fmt.Println("Servidor iniciado com sucesso na porta 8080...")
	// Inicia o servidor
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Erro ao iniciar o servidor: %s\n", err)
	}
}

func getTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	layout := "2006-01-02T15:04:05"
	salaryReceived, _ := time.Parse(layout, "2026-04-26T14:56:05")

	// Criando uma lista de exemplo usando o pacote transaction
	minhasTransacoes := transaction.Transactions{
		transaction.Transaction{
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
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Converte o JSON recebido para a struct
	err = json.Unmarshal(body, &novaTransacao)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Printf("Transação recebida: %+v\n", novaTransacao)
	w.WriteHeader(http.StatusCreated)
}
