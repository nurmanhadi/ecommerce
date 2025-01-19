package controller

import (
	"ecommerce/internal/service"
	"ecommerce/pkg/infrastructure/dto"
	"ecommerce/pkg/infrastructure/response"

	"github.com/gofiber/fiber/v2"
)

type transactionImpl struct {
	transactionService service.TransactionService
}

func NewTransactionController(transactionService *service.TransactionService) TransactionController {
	return &transactionImpl{transactionService: *transactionService}
}
func (h *transactionImpl) AddTrasaction(c *fiber.Ctx) error {
	req := new(dto.TransactionRequestDto)
	if err := c.BodyParser(&req); err != nil {
		return response.ErrorResponse(c, 400, err.Error())
	}
	transaction, err := h.transactionService.AddTransaction(req)
	if err != nil {
		return response.ResponseError(c, err)
	}
	return response.ResponseSuccess(c, 200, transaction)
}
func (h *transactionImpl) CheckTrasaction(c *fiber.Ctx) error {
	orderId := c.Params("orderId")
	status, err := h.transactionService.CheckTransaction(&orderId)
	if err != nil {
		return response.ResponseError(c, err)
	}
	return response.ResponseSuccess(c, 200, status)
}
