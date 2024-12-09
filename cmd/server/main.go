package main

import (
	"transaction/internal/repository"
	"transaction/internal/service"
	"transaction/internal/web"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func main() {
	var logger *zap.Logger
	var db *sqlx.DB
	var app *fiber.App

	repo := repository.CreateNewRepository(logger, db)
	srv := service.CreateNewService(logger, repo)
	wc := web.CreateNewWebController(logger, srv, app)

	wc.RegisterRoutes()
}
