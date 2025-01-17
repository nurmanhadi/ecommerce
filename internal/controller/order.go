package controller

import (
	"ecommerce/internal/service"
	"ecommerce/pkg/infrastructure/dto"
	"ecommerce/pkg/infrastructure/response"

	"github.com/gofiber/fiber/v2"
)

type orderImpl struct {
	orderService service.OrderService
}

func NewOrderController(orderService *service.OrderService) OrderController {
	return &orderImpl{orderService: *orderService}
}
func (h *orderImpl) AddOrder(c *fiber.Ctx) error {
	userId, ok := c.Locals("userId").(string)
	if !ok {
		return response.ErrorResponse(c, 401, "only for user access")
	}
	req := new(dto.OrderAddRequestDto)
	if err := c.BodyParser(&req); err != nil {
		response.ErrorResponse(c, 400, err.Error())
	}
	err := h.orderService.AddOrder(&userId, req)
	if err != nil {
		return response.ResponseError(c, err)
	}
	return response.ResponseSuccess(c, 201, nil)
}
func (h *orderImpl) GetOrders(c *fiber.Ctx) error {
	userId, ok := c.Locals("userId").(string)
	if !ok {
		return response.ErrorResponse(c, 401, "only for user access")
	}
	orders, err := h.orderService.GetOrders(&userId)
	if err != nil {
		return response.ResponseError(c, err)
	}
	return response.ResponseSuccess(c, 200, orders)
}
func (h *orderImpl) GetOrderById(c *fiber.Ctx) error {
	_, ok := c.Locals("userId").(string)
	if !ok {
		return response.ErrorResponse(c, 401, "only for user access")
	}
	orderId := c.Params("orderId")
	order, err := h.orderService.GetOrderById(&orderId)
	if err != nil {
		return response.ResponseError(c, err)
	}
	return response.ResponseSuccess(c, 200, order)
}
func (h *orderImpl) UpdateOrderStatus(c *fiber.Ctx) error {
	orderId := c.Params("orderId")
	status := c.Query("status", "")
	err := h.orderService.UpdateOrderStatus(&orderId, &status)
	if err != nil {
		return response.ResponseError(c, err)
	}
	return response.ResponseSuccess(c, 200, nil)
}
