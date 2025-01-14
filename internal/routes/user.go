package routes

import (
	"ecommerce/internal/controller"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App, c controller.UserController) {
	user := app.Group("/api/v1/users")
	user.Post("/register", c.UserRegister)
	user.Post("/login", c.UserLogin)
}
