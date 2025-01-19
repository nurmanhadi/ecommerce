package repository

import (
	"ecommerce/internal/model"
	"ecommerce/pkg/infrastructure/dto"
	"time"
)

type UserRepository interface {
	RegisterUser(user *model.User) error
	CountUser(email *string) (int, error)
	GetUserById(userId *string) (*model.User, error)
	GetUserByEmail(email *string) (*model.User, error)
	CountUserById(userId *string) (int, error)
}
type ProductRepository interface {
	AddProduct(product *model.Product) (int64, error)
	GetProductById(productId *int) (*model.Product, error)
	UpdateProduct(productId *int, product *dto.ProductUpdateRequestDto) error
	DeleteProduct(productId *int) error
	GetProducts() ([]model.Product, error)
	CountProduct(productId *int) (int, error)
}
type OrderRepository interface {
	AddOrder(order *model.Order) error
	GetOrders(userId *string) ([]model.Order, error)
	GetOrderById(orderId *string) (*model.Order, error)
	UpdateOrderStatus(orderId *string, status *string) error
	CountOrder(orderId *string) (int, error)
}
type PaymentRepository interface {
	AddPayment(payment *model.Payment) error
	UpdateStatus(settlementTime *time.Time, status *string, orderId *string) error
	CountPayment(orderId *string) (int, error)
	GetPaymentByOrderId(orderId *string) (*model.Payment, error)
}
