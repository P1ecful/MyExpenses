package service

import (
	"context"
	"transaction/internal/repository"
	"transaction/internal/web/requests"

	"go.uber.org/zap"
)

type Transactions interface {
	AddExpense(req *requests.AddExpenseRequest) *requests.Response
	Transactions(id int) *requests.TransactionsResponse
	GetBalance(id int) *requests.BalanceResponse
}

type Service struct {
	logger *zap.Logger
	repo   repository.Repository
}

func CreateNewService(log *zap.Logger, repo repository.Repository) *Service {
	return &Service{
		logger: log,
		repo:   repo,
	}
}

// AddExpense method to add user's transaction(expense or income)
func (s *Service) AddExpense(req *requests.AddExpenseRequest) *requests.Response {
	TransID := s.repo.AddTransaction(context.Background(), req)

	return &requests.Response{
		TransactionID: TransID,
		Message:       "Successful",
	}
}

// GetBalance method to check user's balance
func (s *Service) GetBalance(id int) *requests.BalanceResponse {
	balance := s.repo.CheckBalance(context.Background(), id)

	return &requests.BalanceResponse{
		Currency: "USD",
		Balance:  balance,
	}
}

// Transactions method to get all user's transactions
func (s *Service) Transactions(id int) *requests.TransactionsResponse {
	response := s.repo.CheckTransactions(context.Background(), id)

	return &requests.TransactionsResponse{
		Transaction: response,
	}
}
