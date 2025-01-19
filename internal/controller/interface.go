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
type OrderController interface {
	AddOrder(c *fiber.Ctx) error
	GetOrders(c *fiber.Ctx) error
	GetOrderById(c *fiber.Ctx) error
	UpdateOrderStatus(c *fiber.Ctx) error
}
type TransactionController interface {
	CheckTrasaction(c *fiber.Ctx) error
	AddTrasaction(c *fiber.Ctx) error
}
type NotificationController interface {
	AddPaymentOrUpdateStatus(c *fiber.Ctx) error
}
type PaymentController interface {
	GetPaymentByOrderId(c *fiber.Ctx) error
}
