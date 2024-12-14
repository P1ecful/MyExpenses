package service

import (
	"context"
	"transaction/internal/repository"
	"transaction/internal/web/requests"

	"go.uber.org/zap"
)

type Service interface {
	AddExpense(req *requests.AddExpenseRequest) *requests.Response
	Transactions(id int) *requests.TransactionsResponse
	GetBalance(id int) *requests.BalanceResponse
}

type service struct {
	logger *zap.Logger
	repo   repository.Repository
}

func CreateNewService(log *zap.Logger, repo repository.Repository) *service {
	return &service{
		logger: log,
		repo:   repo,
	}
}

// AddExpense is method to add user's transaction(expense or income)
func (s *service) AddExpense(req *requests.AddExpenseRequest) *requests.Response {
	trans_id := s.repo.AddTransaction(context.Background(), req)

	return &requests.Response{
		TransactionID: trans_id,
		Message:       "Succesful",
	}
}

// GetBalance method to check user's balance
func (s *service) GetBalance(id int) *requests.BalanceResponse {
	balance := s.repo.CheckBalance(context.Background(), id)

	return &requests.BalanceResponse{
		Currency: "USD",
		Balance:  balance,
	}
}

// Transactions is method to get all user's transactions
func (s *service) Transactions(id int) *requests.TransactionsResponse {
	response := s.repo.CheckTransactions(context.Background(), id)

	return &requests.TransactionsResponse{
		Transaction: response,
	}
}
