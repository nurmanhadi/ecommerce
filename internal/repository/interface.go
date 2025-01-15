package repository

import (
	"ecommerce/internal/model"
	"ecommerce/pkg/infrastructure/dto"
)

type UserRepository interface {
	RegisterUser(user *model.User) error
	CountUser(email *string) (int, error)
	GetUserById(userId *string) (*model.User, error)
	GetUserByEmail(email *string) (*model.User, error)
}
type ProductRepository interface {
	AddProduct(product *model.Product) error
	GetProductById(productId *int) (*model.Product, error)
	UpdateProduct(productId *int, product *dto.ProductUpdateRequestDto) error
	DeleteProduct(productId *int) error
	GetProducts() ([]model.Product, error)
	CountProduct(productId *int) (int, error)
}
