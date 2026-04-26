package transaction

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ricardovalehrio/dio-expert-session-finance/model/transaction"
	"github.com/ricardovalehrio/dio-expert-session-finance/util"
)

// GetTransactions blablabla
func GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// Use o nome do pacote (transaction) antes da Struct (Transactions)
	var transactions = transaction.Transactions{
		transaction.Transaction{
			Title:     "Salário",
			Amount:    1200.0,
			Type:      0,
			CreatedAt: util.StringToTime("2026-04-26T18:00:05"),
		},
	}

	_ = json.NewEncoder(w).Encode(transactions)
}

// CreateATransactions blablabla
func CreateATransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var res = transaction.Transactions{}
	var body, _ = ioutil.ReadAll(r.Body)

	_ = json.Unmarshal(body, &res)

	fmt.Print(res)

}
