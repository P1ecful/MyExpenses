package requests

import "transaction/internal/models"

type AddExpenseRequest struct {
	Amount   int `json:"amount"`
	Currency int `json:"currency"`
	Category int `json:"category"`
	Type     int `json:"type"`
}

type TransactionsResponse struct {
	Transaction *models.TransactionModel `json:"transaction"`
}

type ExchangeRatesResponse struct {
	BaseCurrency string             `json:"base_currency"`
	Rates        map[string]float64 `json:"rates"`
}

type BalanceResponse struct {
	Currency int     `json:"currency"`
	Balance  float64 `json:"balance"`
}

type Response struct {
	TransactionID int    `json:"transaction_id"`
	Message       string `json:"message"`
}
