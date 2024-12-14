package requests

import "transaction/internal/models"

// AddExpenseRequest is struct of request to add transaction
type AddExpenseRequest struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
	Category string  `json:"category"`
	Type     string  `json:"type"`
}

// UserIdRequest is struct of request for user's methods
type UserIdRequest struct {
	UserId int `json:"user_id"`
}

// TransactionsResponse is response struct of getting transaction history
type TransactionsResponse struct {
	Transaction []models.TransactionModel `json:"transaction"`
}

// ExchangeRatesResponse is response struct of exchange rates
type ExchangeRatesResponse struct {
	BaseCurrency string             `json:"base_currency"`
	Rates        map[string]float64 `json:"rates"`
}

// BalanceResponse is response struct of getting users's balance
type BalanceResponse struct {
	Currency string  `json:"currency"`
	Balance  float64 `json:"balance"`
}

// Response is default response
type Response struct {
	TransactionID int    `json:"transaction_id"`
	Message       string `json:"message"`
}
