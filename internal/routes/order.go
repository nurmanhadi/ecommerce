package routes

import (
	"ecommerce/internal/controller"
	"ecommerce/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func OrderRoute(app *fiber.App, c controller.OrderController) {
	order := app.Group("/api/v1/orders", middleware.AuthMiddleware)
	order.Post("/", c.AddOrder)
	order.Get("/", c.GetOrders)
	order.Get("/:orderId", c.GetOrderById)
	order.Put("/:orderId", c.UpdateOrderStatus)

}
