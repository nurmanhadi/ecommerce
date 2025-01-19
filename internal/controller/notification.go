package controller

import (
	"ecommerce/internal/service"
	"ecommerce/pkg/infrastructure/dto"
	"ecommerce/pkg/infrastructure/response"
	"log"

	"github.com/gofiber/fiber/v2"
)

type notificationImpl struct {
	paymentService service.PaymentService
}

func NewNotificationCOntroller(paymentService *service.PaymentService) NotificationController {
	return &notificationImpl{paymentService: *paymentService}
}
func (h *notificationImpl) AddPaymentOrUpdateStatus(c *fiber.Ctx) error {
	req := new(dto.PaymentRequestDto)
	if err := c.BodyParser(&req); err != nil {
		log.Println(err.Error())
		response.ErrorResponse(c, 400, err.Error())
	}
	err := h.paymentService.AddPaymentOrStatus(req)
	if err != nil {
		return response.ResponseError(c, err)
	}
	return response.ResponseSuccess(c, 200, nil)
}
