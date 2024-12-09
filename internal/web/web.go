package web

import (
	"transaction/internal/service"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type webcontroller struct {
	log *zap.Logger
	app *fiber.App
	srv service.Service
}

func CreateNewWebController(log *zap.Logger, srv service.Service, app *fiber.App) *webcontroller {
	return &webcontroller{
		log: log,
		app: app,
		srv: srv,
	}
}

func (wc *webcontroller) RegisterRoutes() {
	wc.app.Post("/transaction", func(c *fiber.Ctx) error {
		var req service.AddExpenseRequest
		if err := c.BodyParser(&req); err != nil {
			panic(err)
		}

		return c.JSON(wc.srv.AddExpense(&req))
	})

	wc.app.Get("/transactions", func(c *fiber.Ctx) error {
		var req service.IdRequest
		if err := c.BodyParser(&req); err != nil {
			panic(err)
		}

		return c.JSON(wc.srv.Transactions(&req))
	})

	wc.app.Get("/balance", func(c *fiber.Ctx) error {
		var req service.IdRequest
		if err := c.BodyParser(&req); err != nil {
			panic(err)
		}

		return c.JSON(wc.srv.GetBalance(&req))
	})

	wc.app.Get("/exchange-rates", func(c *fiber.Ctx) error {
		return c.JSON(wc.srv.ExchangeRates())
	})
}
