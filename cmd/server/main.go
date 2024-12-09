package main

import (
	"transaction/internal/config"
	"transaction/internal/repository"
	"transaction/internal/service"
	"transaction/internal/web"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

func main() {
	var logger *zap.Logger
	var db *pgxpool.Pool
	var db_cfg *config.PSQLConnection
	var app *fiber.App

	repo := repository.CreatePostgresRepository(logger, db, db_cfg)
	srv := service.CreateNewService(logger, repo)
	wc := web.CreateNewWebController(logger, srv, app)

	wc.RegisterRoutes()
}
