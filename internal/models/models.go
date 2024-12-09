package models

import "time"

type TransactionModel struct {
	UserID        int        `json:"user_id"`
	TransactionID int        `json:"transaction_id"`
	Amount        float64    `json:"amount"`
	Currency      string     `json:"currency"`
	Category      string     `json:"category"`
	Type          string     `json:"type"`
	Date          *time.Time `json:"created_at"`
}
