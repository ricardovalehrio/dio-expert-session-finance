package transaction

import "time"

// Transaction blablabla
type Transaction struct {
	Title     string    `json:"title"`
	Amount    float64   `json:"amount"`
	Type      int       `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

// Transactions blablabla
type Transactions []Transaction
