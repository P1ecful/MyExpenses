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
	CheckTransactions(ctx context.Context, user_id int) []models.TransactionModel
	CheckBalance(ctx context.Context, user_id int) float64
	Disconnect()
}

type pgxstorage struct {
	logger *zap.Logger
	config *config.PSQLConnection
	pool   *pgxpool.Pool
}

// CreatePGXConnection is method to create connetion and database image
func CreatePGXConnection(logger *zap.Logger,
	config *config.PSQLConnection) *pgxstorage {

	pool, err := pgxpool.New(context.Background(),
		fmt.Sprintf(
			"postgresql://%s:%s@%s:%s/%s",
			config.Username, config.Password, config.Host,
			config.Port, config.Database,
		))

	if err != nil {
		logger.Fatal("unable to create connection pool",
			zap.Field(zap.Error(err)))
	}

	if err := pool.Ping(context.Background()); err != nil {
		logger.Fatal("failed to ping database",
			zap.Field(zap.Error(err)))
	}

	logger.Debug("Database connected")

	return &pgxstorage{
		logger: logger,
		config: config,
		pool:   pool,
	}
}

// AddTransaction is method for add transaction in database
func (p *pgxstorage) AddTransaction(ctx context.Context, req *requests.AddExpenseRequest) int {
	rand.Seed(uint64(time.Now().Unix())) // init random seed
	trans_id := rand.Intn(999999)

	query := `insert into transactions 
	(id, user_id, amount, currency, category, type) values
	(@id, @user_id, @amount, @currency, @category, @type)`

	// user_id is const. All time user_id = 1
	args := pgx.NamedArgs{
		"id":       trans_id,
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
			zap.Field(zap.Error(err)))
	}

	p.logger.Debug(fmt.Sprintf("Succesful insert %d", trans_id))
	return trans_id
}

// CheckTransactions is method to get all transactions
func (p *pgxstorage) CheckTransactions(ctx context.Context, user_id int) []models.TransactionModel {
	var transactions []models.TransactionModel
	query := `select * from transactions where user_id = $1`

	rows, err := p.pool.Query(ctx, query, user_id)

	if err != nil {
		p.logger.Debug("Failed get transactions",
			zap.Field(zap.Error(err)))
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
				zap.Field(zap.Error(err)))
		}

		transactions = append(transactions, transaction)
	}

	p.logger.Debug(fmt.Sprintf("User: %d checked transaction history", user_id))
	return transactions
}

// CheckBalance is method for checking balance by user_id
func (p *pgxstorage) CheckBalance(ctx context.Context, user_id int) float64 {
	var balance float64
	query := `select amount from transactions where user_id = $1 and type = 'income'`

	rows, err := p.pool.Query(ctx, query, user_id)

	if err != nil {
		p.logger.Debug("Failed get transactions",
			zap.Field(zap.Error(err)))
	}

	defer rows.Close()

	for rows.Next() {
		var income float64

		if err := rows.Scan(&income); err != nil {
			p.logger.Debug("Failed scan transactions",
				zap.Field(zap.Error(err)))
		}

		balance += income
	}

	p.logger.Debug(fmt.Sprintf("User: %d Checked balance", user_id))
	return balance
}

// Disconnect is method to close database connection
func (p *pgxstorage) Disconnect() {
	p.pool.Close()
}
