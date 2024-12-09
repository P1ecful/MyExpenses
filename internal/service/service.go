package service

import (
	"transaction/internal/models"
	"transaction/internal/repository"

	"go.uber.org/zap"
)

type Service interface {
	AddExpense(req *AddExpenseRequest) *Response
	ExchangeRates() *ExchangeRatesResponse
	Transactions(req *IdRequest) *TransactionsResponse
	GetBalance(req *IdRequest) *BalanceResponse
}

type service struct {
	logger     *zap.Logger
	repository repository.Repository
}

func CreateNewService(log *zap.Logger, repo repository.Repository) *service {
	return &service{
		logger:     log,
		repository: repo,
	}
}

type AddExpenseRequest struct {
	Amount   int `json:"amount"`
	Currency int `json:"currency"`
	Category int `json:"category"`
	Type     int `json:"type"`
}

type IdRequest struct {
	UserID int `jsob:"user_id"`
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

func (s *service) AddExpense(req *AddExpenseRequest) *Response {
	return &Response{}
}

func (s *service) ExchangeRates() *ExchangeRatesResponse {
	return &ExchangeRatesResponse{}
}

func (s *service) GetBalance(req *IdRequest) *BalanceResponse {
	return &BalanceResponse{}
}

func (s *service) Transactions(req *IdRequest) *TransactionsResponse {
	return &TransactionsResponse{}
}
