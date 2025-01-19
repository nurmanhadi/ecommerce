package routes

import (
	"ecommerce/internal/controller"
	"ecommerce/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func PaymentRoute(app *fiber.App, c controller.PaymentController) {
	payment := app.Group("api/v1/orders/:orderId/payments", middleware.AuthMiddleware)
	payment.Get("/", c.GetPaymentByOrderId)
}
