package controller

import (
	"ecommerce/internal/service"
	"ecommerce/pkg/infrastructure/dto"
	"ecommerce/pkg/infrastructure/response"

	"github.com/gofiber/fiber/v2"
)

type productImpl struct {
	productService service.ProductService
}

func NewProductController(productService *service.ProductService) ProductController {
	return &productImpl{productService: *productService}
}
func (h *productImpl) GetProducts(c *fiber.Ctx) error {
	products, err := h.productService.GetProducts()
	if err != nil {
		return response.ResponseError(c, err)
	}
	return response.ResponseSuccess(c, 200, products)
}
func (h *productImpl) GetProductById(c *fiber.Ctx) error {
	productId := c.Params("productId")
	product, err := h.productService.GetProductById(&productId)
	if err != nil {
		return response.ResponseError(c, err)
	}
	return response.ResponseSuccess(c, 200, product)
}
func (h *productImpl) AddProduct(c *fiber.Ctx) error {
	req := new(dto.ProductAddRequestDto)
	if err := c.BodyParser(&req); err != nil {
		return response.ErrorResponse(c, 400, err.Error())
	}
	if err := h.productService.AddProduct(req); err != nil {
		return response.ResponseError(c, err)
	}
	return response.ResponseSuccess(c, 201, nil)
}
func (h *productImpl) UpdateProduct(c *fiber.Ctx) error {
	productId := c.Params("productId")
	req := new(dto.ProductUpdateRequestDto)
	if err := c.BodyParser(&req); err != nil {
		return response.ErrorResponse(c, 400, err.Error())
	}
	if err := h.productService.UpdateProduct(&productId, req); err != nil {
		return response.ResponseError(c, err)
	}
	return response.ResponseSuccess(c, 200, nil)
}
func (h *productImpl) DeleteProduct(c *fiber.Ctx) error {
	productId := c.Params("productId")

	if err := h.productService.DeleteProduct(&productId); err != nil {
		return response.ResponseError(c, err)
	}
	return response.ResponseSuccess(c, 200, nil)
}
