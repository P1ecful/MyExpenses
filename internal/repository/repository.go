package repository

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Repository interface {
	ConnectRepository()
	MigrateRepository()
	AddTransaction()
	CheckTransactions()
	CheckBalance()
}

type repository struct {
	logger     *zap.Logger
	repository *sqlx.DB
}

func CreateNewRepository(log *zap.Logger, repo *sqlx.DB) *repository {
	return &repository{
		logger:     log,
		repository: repo,
	}
}

func (r *repository) ConnectRepository() {}
func (r *repository) MigrateRepository() {}
func (r *repository) AddTransaction()    {}
func (r *repository) CheckTransactions() {}
func (r *repository) CheckBalance()      {}
