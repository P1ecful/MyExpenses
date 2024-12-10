package service

import (
	"net/http"
	"transaction/internal/config"
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
	response, err := http.Get(config.CurrencyURL)

	if err != nil {
		s.logger.Fatal("Failed request to ExchangeRates",
			zap.Field(zap.Error(err)))
	}

	return &requests.Response{
		Message: response.Status,
	}
}

func (s *service) ExchangeRates() *requests.ExchangeRatesResponse {
	rates := make(map[string]float64)

	return &requests.ExchangeRatesResponse{
		BaseCurrency: "USD",
		Rates:        rates,
	}
}

func (s *service) GetBalance(id int) *requests.BalanceResponse {
	return &requests.BalanceResponse{}
}

func (s *service) Transactions(id int) *requests.TransactionsResponse {
	return &requests.TransactionsResponse{}
}
