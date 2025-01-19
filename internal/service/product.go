package service

import (
	"ecommerce/internal/model"
	"ecommerce/internal/repository"
	"ecommerce/pkg/infrastructure/dto"
	"ecommerce/pkg/infrastructure/exception"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
)

type productImpl struct {
	productRepository repository.ProductRepository
	validator         *validator.Validate
}

func NewProductService(productRepository *repository.ProductRepository, validator *validator.Validate) ProductService {
	return &productImpl{productRepository: *productRepository, validator: validator}
}

func (s *productImpl) GetProducts() ([]model.Product, error) {
	products, err := s.productRepository.GetProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}
func (s *productImpl) GetProductById(productId *string) (*model.Product, error) {
	convertId, err := strconv.ParseInt(*productId, 10, 32)
	if err != nil {
		return nil, err
	}
	id := int(convertId)
	product, err := s.productRepository.GetProductById(&id)
	if err != nil {
		return nil, exception.ProductNotFound
	}
	return product, nil
}
func (s *productImpl) AddProduct(req *dto.ProductAddRequestDto) (map[string]interface{}, error) {
	if err := s.validator.Struct(req); err != nil {
		return nil, err
	}
	product := &model.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	id, err := s.productRepository.AddProduct(product)
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"product_id": id,
	}, nil
}
func (s *productImpl) UpdateProduct(productId *string, req *dto.ProductUpdateRequestDto) error {
	convertId, err := strconv.ParseInt(*productId, 10, 32)
	if err != nil {
		return err
	}
	id := int(convertId)
	count, err := s.productRepository.CountProduct(&id)
	if err != nil {
		return err
	}
	if count == 0 {
		return exception.ProductNotFound
	}
	if err := s.validator.Struct(req); err != nil {
		return err
	}
	if err := s.productRepository.UpdateProduct(&id, req); err != nil {
		return err
	}
	return nil
}
func (s *productImpl) DeleteProduct(productId *string) error {
	convertId, err := strconv.ParseInt(*productId, 10, 32)
	if err != nil {
		return err
	}
	id := int(convertId)
	count, err := s.productRepository.CountProduct(&id)
	if err != nil {
		return err
	}
	if count == 0 {
		return exception.ProductNotFound
	}
	if err := s.productRepository.DeleteProduct(&id); err != nil {
		return err
	}
	return nil
}
