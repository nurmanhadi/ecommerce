package unit

import (
	"ecommerce/config"
	gateway "ecommerce/external/midtrans"
	"ecommerce/internal/model"
	"fmt"
	"testing"
	"time"

	"github.com/midtrans/midtrans-go"
)

func TestMidtransCreateTransaction(t *testing.T) {
	config.LoadConfig()
	midtrans.ServerKey = config.Viper.Midtrans.Key
	midtrans.Environment = midtrans.Sandbox
	product := &model.Product{
		Id:          6,
		Name:        "test",
		Description: "test",
		Price:       50000,
		Stock:       100,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	order := &model.Order{
		Id:          "test4",
		UserId:      "user",
		ProductId:   6,
		Quantity:    2,
		GrossAmount: 100000,
		Status:      "pending",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	response, err := gateway.AddTransaction(order, product)
	if err != nil {
		panic(err)
	}
	fmt.Println("status code " + response.StatusCode)
	fmt.Println(response)
}
func TestMidtransCheckTransaction(t *testing.T) {
	config.LoadConfig()
	midtrans.ServerKey = config.Viper.Midtrans.Key
	midtrans.Environment = midtrans.Sandbox
	orderId := "test4"
	response, err := gateway.CheckTransaction(&orderId)
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}
