package repository

import (
	"context"
	"fmt"
	"transaction/internal/config"
	"transaction/internal/models"
	"transaction/internal/web/requests"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Repository interface {
	ConnectRepository() *pgxpool.Pool
	AddTransaction(req *requests.AddExpenseRequest) error
	CheckTransactions(id int) ([]models.TransactionModel, error)
	CheckBalance(id int) (string, float64, error)
	MigrateRepository()
}

type postgres struct {
	logger *zap.Logger
	config *config.PSQLConnection
}

func CreatePostgresRepository(log *zap.Logger, cfg *config.PSQLConnection) *postgres {
	return &postgres{
		logger: log,
		config: cfg,
	}
}

func (p *postgres) ConnectRepository() *pgxpool.Pool {
	pool, err := pgxpool.New(context.Background(),
		fmt.Sprintf(
			"postgresql://%s:%s@%s:%s/%s",
			p.config.Username, p.config.Password, p.config.Host,
			p.config.Port, p.config.Database,
		))

	if err != nil {
		p.logger.Fatal("unable to create connection pool",
			zap.Field(zap.Error(err)))
	}

	if err := pool.Ping(context.Background()); err != nil {
		p.logger.Fatal("failed to ping database",
			zap.Field(zap.Error(err)))
	}

	return pool
}

func (p *postgres) AddTransaction(req *requests.AddExpenseRequest) error {
	return nil
}

func (p *postgres) CheckTransactions(id int) ([]models.TransactionModel, error) {
	return nil, nil
}

func (p *postgres) CheckBalance(id int) (string, float64, error) {
	return "", 1.05, nil
}

func (p *postgres) MigrateRepository() {}
