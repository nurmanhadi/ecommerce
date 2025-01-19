package service

import (
	"ecommerce/internal/model"
	"ecommerce/internal/repository"
	"ecommerce/pkg/infrastructure/dto"
	"ecommerce/pkg/infrastructure/exception"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type paymentImpl struct {
	paymentRepository repository.PaymentRepository
	productRepository repository.ProductRepository
	orderRepository   repository.OrderRepository
	validator         *validator.Validate
}

func NewPaymentService(paymentRepository *repository.PaymentRepository, productRepository *repository.ProductRepository, orderRepository *repository.OrderRepository, validator *validator.Validate) PaymentService {
	return &paymentImpl{paymentRepository: *paymentRepository, productRepository: *productRepository, orderRepository: *orderRepository, validator: validator}
}
func (s *paymentImpl) AddPaymentOrStatus(req *dto.PaymentRequestDto) error {
	if err := s.validator.Struct(req); err != nil {
		return err
	}
	countPyament, err := s.paymentRepository.CountPayment(&req.OrderId)
	if err != nil {
		return err
	}
	if countPyament == 0 {
		traTime, err := time.Parse(time.DateTime, req.TransactionTime)
		if err != nil {
			return err
		}
		payment := &model.Payment{
			Id:                uuid.NewString(),
			OrderId:           req.OrderId,
			TransactionId:     req.TransactionId,
			MerchantId:        req.MerchantId,
			TransactionStatus: req.TransactionStatus,
			SignatureKey:      req.SignatureKey,
			PaymentType:       req.PaymentType,
			GrossAmount:       req.GrossAmount,
			FraudStatus:       req.FraudStatus,
			Currency:          req.Currency,
			TransactionTime:   traTime,
		}
		err = s.paymentRepository.AddPayment(payment)
		if err != nil {
			return err
		}
	} else {
		setTime, err := time.Parse(time.DateTime, req.SettlementTime)
		if err != nil {
			return err
		}
		err = s.paymentRepository.UpdateStatus(&setTime, &req.TransactionStatus, &req.OrderId)
		if err != nil {
			return err
		}
		order, err := s.orderRepository.GetOrderById(&req.OrderId)
		if err != nil {
			return exception.OrderNotFound
		}
		product, err := s.productRepository.GetProductById(&order.ProductId)
		if err != nil {
			return exception.ProductNotFound
		}
		stokUpdate := product.Stock - order.Quantity
		updateStok := &dto.ProductUpdateRequestDto{
			Stock: &stokUpdate,
		}
		err = s.productRepository.UpdateProduct(&product.Id, updateStok)
		if err != nil {
			return err
		}
		err = s.orderRepository.UpdateOrderStatus(&req.OrderId, &req.TransactionStatus)
		if err != nil {
			return err
		}
	}
	return nil
}
func (s *paymentImpl) GetPaymentByOrderId(orderId *string) (*model.Payment, error) {
	payment, err := s.paymentRepository.GetPaymentByOrderId(orderId)
	if err != nil {
		return nil, exception.PaymentNotFound
	}
	return payment, nil
}
