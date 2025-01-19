package controller

import (
	"ecommerce/internal/service"
	"ecommerce/pkg/infrastructure/response"

	"github.com/gofiber/fiber/v2"
)

type paymentImpl struct {
	paymentService service.PaymentService
}

func NewPaymentController(paymentService *service.PaymentService) PaymentController {
	return &paymentImpl{paymentService: *paymentService}
}
func (h *paymentImpl) GetPaymentByOrderId(c *fiber.Ctx) error {
	orderId := c.Params("orderId")
	payment, err := h.paymentService.GetPaymentByOrderId(&orderId)
	if err != nil {
		return response.ResponseError(c, err)
	}
	return response.ResponseSuccess(c, 200, payment)
}
