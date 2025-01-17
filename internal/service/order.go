package service

import (
	"ecommerce/internal/model"
	"ecommerce/internal/repository"
	"ecommerce/pkg/infrastructure/dto"
	"ecommerce/pkg/infrastructure/exception"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type orderImpl struct {
	orderRepository   repository.OrderRepository
	productRepository repository.ProductRepository
	userRepository    repository.UserRepository
	validator         *validator.Validate
}

func NewOrderService(orderRepository *repository.OrderRepository, productRepository *repository.ProductRepository, userRepository *repository.UserRepository, validator *validator.Validate) OrderService {
	return &orderImpl{orderRepository: *orderRepository, productRepository: *productRepository, userRepository: *userRepository, validator: validator}
}
func (s *orderImpl) AddOrder(userId *string, req *dto.OrderAddRequestDto) error {
	countUser, err := s.userRepository.CountUserById(userId)
	if err != nil {
		return err
	}
	if countUser == 0 {
		return exception.UserNotFound
	}
	product, err := s.productRepository.GetProductById(&req.ProductId)
	if err != nil {
		return exception.ProductNotFound
	}
	if err := s.validator.Struct(req); err != nil {
		return err
	}
	orderId := uuid.NewString()
	order := &model.Order{
		Id:          orderId,
		UserId:      *userId,
		ProductId:   product.Id,
		Quantity:    req.Quantity,
		GrossAmount: product.Price * req.Quantity,
		Status:      "pending",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	err = s.orderRepository.AddOrder(order)
	if err != nil {
		return err
	}
	return nil
}
func (s *orderImpl) GetOrders(userId *string) ([]model.Order, error) {
	countUser, err := s.userRepository.CountUserById(userId)
	if err != nil {
		return nil, err
	}
	if countUser == 0 {
		return nil, exception.UserNotFound
	}
	orders, err := s.orderRepository.GetOrders(userId)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
func (s *orderImpl) GetOrderById(orderId *string) (*model.Order, error) {
	order, err := s.orderRepository.GetOrderById(orderId)
	fmt.Println(orderId)
	if err != nil {
		return nil, exception.OrderNotFound
	}
	return order, nil
}
func (s *orderImpl) UpdateOrderStatus(orderId *string, status *string) error {
	countOrder, err := s.orderRepository.CountOrder(orderId)
	if err != nil {
		return err
	}
	if countOrder == 0 {
		return exception.OrderNotFound
	}
	if *status == "" {
		return exception.OrderStatusQueryIsRequired
	}
	if len(*status) > 20 {
		return exception.OrderStatusQueryIsRequired
	}
	err = s.orderRepository.UpdateOrderStatus(orderId, status)
	if err != nil {
		return err
	}
	return nil
}
