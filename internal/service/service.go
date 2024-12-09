package service

import (
	"transaction/internal/repository"
	"transaction/internal/web/requests"

	"go.uber.org/zap"
)

type Service interface {
	AddExpense(req *requests.AddExpenseRequest) *requests.Response
	ExchangeRates() *requests.ExchangeRatesResponse
	Transactions(id int) *requests.TransactionsResponse
	GetBalance(id int) *requests.BalanceResponse
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

func (s *service) AddExpense(req *requests.AddExpenseRequest) *requests.Response {
	return &requests.Response{}
}

func (s *service) ExchangeRates() *requests.ExchangeRatesResponse {
	return &requests.ExchangeRatesResponse{}
}

func (s *service) GetBalance(id int) *requests.BalanceResponse {
	return &requests.BalanceResponse{}
}

func (s *service) Transactions(id int) *requests.TransactionsResponse {
	return &requests.TransactionsResponse{}
}
