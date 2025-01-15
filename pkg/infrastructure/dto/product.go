package dto

type ProductAddRequestDto struct {
	Name        string `json:"name" validate:"required,min=1,max=100"`
	Description string `json:"description" validate:"required,min=1"`
	Price       int    `json:"price" validate:"required,gte=0,lte=1000000000"`
	Stock       int    `json:"stock" validate:"required,gte=0,lte=1000000000"`
}
type ProductUpdateRequestDto struct {
	Name        *string `json:"name" validate:"omitempty,min=1,max=100"`
	Description *string `json:"description" validate:"omitempty,min=1"`
	Price       *int    `json:"price" validate:"omitempty,gte=0,lte=1000000000"`
	Stock       *int    `json:"stock" validate:"omitempty,gte=0,lte=1000000000"`
}
