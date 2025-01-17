package dto

type OrderAddRequestDto struct {
	ProductId int `json:"product_id" validate:"required"`
	Quantity  int `json:"quantity" validate:"required,gte=0,lte=1000000"`
}
