package repository

import (
	"context"
	"fmt"
	"time"
	"transaction/internal/config"
	"transaction/internal/models"
	"transaction/internal/web/requests"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"golang.org/x/exp/rand"
)

type Repository interface {
	AddTransaction(ctx context.Context, req *requests.AddExpenseRequest) int
	CheckTransactions(ctx context.Context, UserID int) []models.TransactionModel
	CheckBalance(ctx context.Context, UserID int) float64
	Disconnect()
}

type Pgxstorage struct {
	logger *zap.Logger
	config *config.PSQLConnection
	pool   *pgxpool.Pool
}

// CreatePGXConnection is method to create connection and database image
func CreatePGXConnection(logger *zap.Logger,
	config *config.PSQLConnection) *Pgxstorage {

	pool, err := pgxpool.New(context.Background(),
		fmt.Sprintf(
			"postgresql://%s:%s@%s:%s/%s",
			config.Username, config.Password, config.Host,
			config.Port, config.Database,
		))

	if err != nil {
		logger.Fatal("unable to create connection pool",
			zap.Error(err))
	}

	if err := pool.Ping(context.Background()); err != nil {
		logger.Fatal("failed to ping database",
			zap.Error(err))
	}

	logger.Debug("Database connected")

	return &Pgxstorage{
		logger: logger,
		config: config,
		pool:   pool,
	}
}

// AddTransaction is method for add transaction in database
func (p *Pgxstorage) AddTransaction(ctx context.Context, req *requests.AddExpenseRequest) int {
	rand.Seed(uint64(time.Now().Unix())) // init random seed
	TransID := rand.Intn(999999)

	query := `insert into transactions 
	(id, user_id, amount, currency, category, type) values
	(@id, @user_id, @amount, @currency, @category, @type)`

	// user_id is const. All time user_id = 1
	args := pgx.NamedArgs{
		"id":       TransID,
		"user_id":  1,
		"amount":   req.Amount,
		"currency": req.Currency,
		"category": req.Category,
		"type":     req.Type,
	}

	// Insert into database
	_, err := p.pool.Exec(ctx, query, args)
	if err != nil {
		p.logger.Fatal("Failed insert into database",
			zap.Error(err))
	}

	p.logger.Debug(fmt.Sprintf("Succesful insert %d", TransID))
	return TransID
}

// CheckTransactions is method to get all transactions
func (p *Pgxstorage) CheckTransactions(ctx context.Context, UserID int) []models.TransactionModel {
	var transactions []models.TransactionModel
	query := `select * from transactions where user_id = $1`

	rows, err := p.pool.Query(ctx, query, UserID)

	if err != nil {
		p.logger.Debug("Failed get transactions",
			zap.Error(err))
	}

	defer rows.Close()

	for rows.Next() {
		var transaction models.TransactionModel
		if err := rows.Scan(
			&transaction.UserID,
			&transaction.TransactionID,
			&transaction.Amount,
			&transaction.Currency,
			&transaction.Category,
			&transaction.Category,
			&transaction.Date); err != nil {

			p.logger.Debug("Failed scan transactions",
				zap.Error(err))
		}

		transactions = append(transactions, transaction)
	}

	p.logger.Debug(fmt.Sprintf("User: %d checked transaction history", UserID))
	return transactions
}

// CheckBalance is method for checking balance by user_id
func (p *Pgxstorage) CheckBalance(ctx context.Context, UserID int) float64 {
	var balance float64
	query := `select amount from transactions where user_id = $1 and type = 'income'`

	rows, err := p.pool.Query(ctx, query, UserID)

	if err != nil {
		p.logger.Debug("Failed get transactions",
			zap.Error(err))
	}

	defer rows.Close()

	for rows.Next() {
		var income float64

		if err := rows.Scan(&income); err != nil {
			p.logger.Debug("Failed scan transactions",
				zap.Error(err))
		}

		balance += income
	}

	p.logger.Debug(fmt.Sprintf("User: %d Checked balance", UserID))
	return balance
}

// Disconnect is method to close database connection
func (p *Pgxstorage) Disconnect() {
	p.pool.Close()
}
