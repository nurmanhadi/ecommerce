package controller

import (
	"ecommerce/internal/service"
	"ecommerce/pkg/infrastructure/dto"
	"ecommerce/pkg/infrastructure/response"

	"github.com/gofiber/fiber/v2"
)

type userImpl struct {
	userService service.UserService
}

func NewUserController(userService *service.UserService) UserController {
	return &userImpl{userService: *userService}
}
func (h *userImpl) UserRegister(c *fiber.Ctx) error {
	req := new(dto.UserRegisterRequestDto)
	if err := c.BodyParser(req); err != nil {
		return response.ErrorResponse(c, 400, err.Error())
	}
	err := h.userService.UserRegister(req)
	if err != nil {
		return response.ResponseError(c, err)
	}
	return response.ResponseSuccess(c, 201, nil)
}
func (h *userImpl) UserLogin(c *fiber.Ctx) error {
	req := new(dto.UserLoginRequestDto)
	if err := c.BodyParser(req); err != nil {
		return response.ErrorResponse(c, 400, err.Error())
	}
	token, err := h.userService.UserLogin(req)
	if err != nil {
		return response.ResponseError(c, err)
	}
	return response.ResponseSuccess(c, 200, token)
}
