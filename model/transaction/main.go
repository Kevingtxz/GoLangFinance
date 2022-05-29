package transaction

import "time"

type Transaction struct {
	Title    string    `json:"title"`
	Amount   float32   `json:"amount"`
	Type     int       `json:"type"`
	CreateAt time.Time `json:"create_at"`
}

type Transactions []Transaction
