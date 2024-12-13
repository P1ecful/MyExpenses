package web

import (
	"transaction/internal/service"
	"transaction/internal/web/requests"

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
		var req requests.AddExpenseRequest
		if err := c.BodyParser(&req); err != nil {
			wc.log.Debug("Failed path: /transaction",
				zap.Field(zap.Error(err)))
		}

		return c.JSON(wc.srv.AddExpense(&req))
	})

	wc.app.Post("/transactions", func(c *fiber.Ctx) error {
		var req requests.UserIdRequest
		if err := c.BodyParser(&req); err != nil {
			wc.log.Debug("Failed path: /transactions",
				zap.Field(zap.Error(err)))
		}

		return c.JSON(wc.srv.Transactions(req.UserId))
	})

	wc.app.Post("/balance", func(c *fiber.Ctx) error {
		var req int
		if err := c.BodyParser(&req); err != nil {
			wc.log.Debug("Failed path: /balance",
				zap.Field(zap.Error(err)))
		}

		return c.JSON(wc.srv.GetBalance(req))
	})

	wc.app.Get("/exchange-rates", func(c *fiber.Ctx) error {
		return c.JSON(wc.srv.ExchangeRates())
	})
}
