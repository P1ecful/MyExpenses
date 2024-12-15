package web

import (
	"transaction/internal/service"
	"transaction/internal/web/requests"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Controller struct {
	log *zap.Logger
	app *fiber.App
	srv service.Transactions
}

func CreateNewWebController(log *zap.Logger, srv service.Transactions, app *fiber.App) *Controller {
	return &Controller{
		log: log,
		app: app,
		srv: srv,
	}
}

func (wc *Controller) RegisterRoutes() {
	wc.app.Post("/transaction", func(c *fiber.Ctx) error {
		var req requests.AddExpenseRequest
		if err := c.BodyParser(&req); err != nil {
			wc.log.Debug("Failed path: /transaction", zap.Error(err))
		}

		return c.JSON(wc.srv.AddExpense(&req))
	})

	wc.app.Post("/transactions", func(c *fiber.Ctx) error {
		var req requests.UserIdRequest
		if err := c.BodyParser(&req); err != nil {
			wc.log.Debug("Failed path: /transactions", zap.Error(err))
		}

		return c.JSON(wc.srv.Transactions(req.UserId))
	})

	wc.app.Post("/balance", func(c *fiber.Ctx) error {
		var req requests.UserIdRequest
		if err := c.BodyParser(&req); err != nil {
			wc.log.Debug("Failed path: /balance", zap.Error(err))
		}

		return c.JSON(wc.srv.GetBalance(req.UserId))
	})
}
