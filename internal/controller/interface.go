package controller

import "github.com/gofiber/fiber/v2"

type UserController interface {
	UserRegister(c *fiber.Ctx) error
	UserLogin(c *fiber.Ctx) error
}
type ProductController interface {
	GetProducts(c *fiber.Ctx) error
	GetProductById(c *fiber.Ctx) error
	AddProduct(c *fiber.Ctx) error
	UpdateProduct(c *fiber.Ctx) error
	DeleteProduct(c *fiber.Ctx) error
}
