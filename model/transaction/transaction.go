// Arquivo: model/transaction/transaction.go
package transaction

import "time"

// Transaction representa uma única transação financeira
type Transaction struct {
	Title     string    `json:"title"`
	Amount    float64   `json:"amount"`
	Type      int       `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

// Transactions é um slice (lista) de transações
type Transactions []Transaction
