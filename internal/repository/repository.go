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
	ConnectRepository(ctx context.Context) error
	AddTransaction(ctx context.Context, t *requests.AddExpenseRequest) error
	CheckTransactions(ctx context.Context, id int) ([]models.TransactionModel, error)
	CheckBalance(ctx context.Context, id int) (string, float64, error)
	MigrateRepository()
	CloseRepository()
}

type postgres struct {
	logger *zap.Logger
	repo   *pgxpool.Pool
	config *config.PSQLConnection
}

func CreatePostgresRepository(log *zap.Logger, repo *pgxpool.Pool, cfg *config.PSQLConnection) *postgres {
	return &postgres{
		logger: log,
		repo:   repo,
		config: cfg,
	}
}

func (p *postgres) ConnectRepository(ctx context.Context) error {
	pool, err := pgxpool.New(ctx,
		fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			p.config.Host, p.config.Port, p.config.Username,
			p.config.Password, p.config.Database,
		))

	if err != nil {
		p.logger.Debug("unable to create connection pool",
			zap.Field(zap.Error(err)),
		)

	}

	if err := pool.Ping(ctx); err != nil {
		p.logger.Debug("failes to ping database",
			zap.Field(zap.Error(err)))
	}

	return nil
}

func (p *postgres) AddTransaction(ctx context.Context, t *requests.AddExpenseRequest) error {
	query := `insert into ExchangeCur () values ()`
	fmt.Println(query)

	return nil
}

func (p *postgres) CheckTransactions(ctx context.Context, id int) ([]models.TransactionModel, error) {
	transactions := []models.TransactionModel{}

	return transactions, nil
}

func (p *postgres) CheckBalance(ctx context.Context, id int) (string, float64, error) {
	cur := "USD"
	bal := 10.65

	return cur, bal, nil
}

func (p *postgres) CloseRepository() {
	p.repo.Close()
}

func (p *postgres) MigrateRepository() {}
