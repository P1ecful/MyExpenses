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
	if err := godotenv.Load("config/local.env"); err != nil {
		panic(err)
	}
}

func main() {
	var logger, _ = zap.NewDevelopment() // logger init
	quit := make(chan os.Signal, 1)      // creating quit goroutine

	postgres := repository.CreatePostgresRepository(logger, &config.PSQLConnection{
		Host:     os.Getenv("DATABASE_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Database: os.Getenv("POSTGRES_DB"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Username: os.Getenv("POSTGRES_USER"),
	})

	repo := postgres.ConnectRepository() // connecting database
	logger.Info("Database connected")

	wApp := fiber.New()                                 // fiber init
	srv := service.CreateNewService(logger, postgres)   // service init
	wc := web.CreateNewWebController(logger, srv, wApp) // web controller init
	wc.RegisterRoutes()                                 // registering handlers

	// start service and graceful shutdown
	go func() {
		if err := wApp.Listen(os.Getenv("SERVICE_PORT")); err != nil {
			signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
			<-quit

			logger.Fatal("Can`t shutdown service",
				zap.Field(zap.Error(err)))
		}
	}()

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	repo.Close() // close database
	logger.Info("Database disconnected")
	defer cancel()

	if err := wApp.Shutdown(); err != nil { // try to stop server
		logger.Info("Failed to stop server")

		return
	}

	logger.Info("Server stopped")
}
