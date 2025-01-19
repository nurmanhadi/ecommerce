package service

import (
	"ecommerce/external/midtrans"
	"ecommerce/internal/repository"
	"ecommerce/pkg/infrastructure/dto"
	"ecommerce/pkg/infrastructure/exception"

	"github.com/go-playground/validator/v10"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

type transactionImpl struct {
	orderRepository   repository.OrderRepository
	productRepository repository.ProductRepository
	validator         *validator.Validate
}

func NewTransactionService(orderRepository *repository.OrderRepository, productRepository *repository.ProductRepository, validator *validator.Validate) TransactionService {
	return &transactionImpl{orderRepository: *orderRepository, productRepository: *productRepository, validator: validator}
}
func (s *transactionImpl) AddTransaction(req *dto.TransactionRequestDto) (*snap.Response, error) {
	if err := s.validator.Struct(req); err != nil {
		return nil, err
	}
	order, err := s.orderRepository.GetOrderById(&req.OrderId)
	if err != nil {
		return nil, exception.OrderNotFound
	}
	product, err := s.productRepository.GetProductById(&order.ProductId)
	if err != nil {
		return nil, exception.ProductNotFound
	}
	midtrans, err := midtrans.AddTransaction(order, product)
	if err != nil {
		return nil, err
	}
	return midtrans, nil
}
func (s *transactionImpl) CheckTransaction(orderId *string) (*coreapi.TransactionStatusResponse, error) {
	countOrder, err := s.orderRepository.CountOrder(orderId)
	if err != nil {
		return nil, err
	}
	if countOrder == 0 {
		return nil, exception.OrderNotFound
	}
	status, err := midtrans.CheckTransaction(orderId)
	if err != nil {
		return nil, err
	}
	return status, nil
}
