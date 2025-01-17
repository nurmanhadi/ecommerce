package model

import "time"

type Order struct {
	Id          string    `json:"id"`
	UserId      string    `json:"user_id"`
	ProductId   int       `json:"product_id"`
	Quantity    int       `json:"quantity"`
	GrossAmount int       `json:"gross_amount"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
