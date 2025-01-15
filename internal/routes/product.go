package routes

import (
	"ecommerce/internal/controller"

	"github.com/gofiber/fiber/v2"
)

func ProductRoute(app *fiber.App, c controller.ProductController) {
	product := app.Group("/api/v1/products")
	product.Get("/", c.GetProducts)
	product.Get("/:productId", c.GetProductById)
	product.Post("/", c.AddProduct)
	product.Patch("/:productId", c.UpdateProduct)
	product.Delete("/:productId", c.DeleteProduct)
}
