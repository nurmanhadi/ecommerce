package service

import (
	"ecommerce/internal/model"
	"ecommerce/pkg/infrastructure/dto"
)

type UserService interface {
	UserRegister(req *dto.UserRegisterRequestDto) error
	UserLogin(req *dto.UserLoginRequestDto) (map[string]interface{}, error)
}
type ProductService interface {
	GetProducts() ([]model.Product, error)
	GetProductById(productId *string) (*model.Product, error)
	AddProduct(req *dto.ProductAddRequestDto) error
	UpdateProduct(productId *string, req *dto.ProductUpdateRequestDto) error
	DeleteProduct(productId *string) error
}
