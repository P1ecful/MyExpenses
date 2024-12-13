package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
	"transaction/internal/config"
	"transaction/internal/repository"
	"transaction/internal/service"
	"transaction/internal/web"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func init() {
	// loading env file with service, database configuration
	if err := godotenv.Load("config/local.env"); err != nil {
		panic(err)
	}
}

func main() {
	var logger, _ = zap.NewDevelopment() // logger init

	postgres := repository.CreatePGXConnection(logger, &config.PSQLConnection{
		Host:     os.Getenv("DATABASE_HOST"),
		Port:     os.Getenv("DATABASE_PORT"),
		Database: os.Getenv("DATABASE_TABLE"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		Username: os.Getenv("DATABASE_USER"),
	})

	wApp := fiber.New()                                 // fiber init
	srv := service.CreateNewService(logger, postgres)   // service init
	wc := web.CreateNewWebController(logger, srv, wApp) // web controller init
	wc.RegisterRoutes()                                 // registering handlers

	quit := make(chan os.Signal, 1) // creating quit goroutine

	// start service and graceful shutdown
	go func() {
		if err := wApp.Listen(os.Getenv("SERVICE_PORT")); err != nil {
			logger.Fatal("Can't shutdown service",
				zap.Field(zap.Error(err)))
		}
	}()

	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	postgres.Disconnect() // close database
	logger.Debug("Database disconnected")
	defer cancel()

	if err := wApp.Shutdown(); err != nil { // try to stop server
		logger.Info("Failed to stop server")

		return
	}

	logger.Debug("Server stopped")
}
