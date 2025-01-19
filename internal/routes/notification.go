package routes

import (
	"ecommerce/internal/controller"

	"github.com/gofiber/fiber/v2"
)

func NotificationRoute(app *fiber.App, c controller.NotificationController) {
	notif := app.Group("/api/v1/notifications")
	notif.Post("/payments", c.AddPaymentOrUpdateStatus)
}
