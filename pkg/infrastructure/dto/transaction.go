package dto

type TransactionRequestDto struct {
	OrderId string `json:"order_id" validate:"required,max=36"`
}
