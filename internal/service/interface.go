package service

import (
	"ecommerce/internal/model"
	"ecommerce/pkg/infrastructure/dto"

	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

type UserService interface {
	UserRegister(req *dto.UserRegisterRequestDto) error
	UserLogin(req *dto.UserLoginRequestDto) (map[string]interface{}, error)
}
type ProductService interface {
	GetProducts() ([]model.Product, error)
	GetProductById(productId *string) (*model.Product, error)
	AddProduct(req *dto.ProductAddRequestDto) (map[string]interface{}, error)
	UpdateProduct(productId *string, req *dto.ProductUpdateRequestDto) error
	DeleteProduct(productId *string) error
}
type OrderService interface {
	AddOrder(userId *string, req *dto.OrderAddRequestDto) (map[string]interface{}, error)
	GetOrders(userId *string) ([]model.Order, error)
	GetOrderById(orderId *string) (*model.Order, error)
	UpdateOrderStatus(orderId *string, status *string) error
}
type TransactionService interface {
	AddTransaction(req *dto.TransactionRequestDto) (*snap.Response, error)
	CheckTransaction(orderId *string) (*coreapi.TransactionStatusResponse, error)
}
type PaymentService interface {
	AddPaymentOrStatus(req *dto.PaymentRequestDto) error
	GetPaymentByOrderId(orderId *string) (*model.Payment, error)
}
