package controller

import "github.com/gofiber/fiber/v2"

type UserController interface {
	UserRegister(c *fiber.Ctx) error
	UserLogin(c *fiber.Ctx) error
}
