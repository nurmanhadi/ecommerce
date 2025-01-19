package routes

import (
	"ecommerce/internal/controller"
	"ecommerce/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func TransactionRoute(app *fiber.App, c controller.TransactionController) {
	transaction := app.Group("/api/v1/transactions", middleware.AuthMiddleware)
	transaction.Post("/", c.AddTrasaction)
	transaction.Get("/:orderId", c.CheckTrasaction)
}
