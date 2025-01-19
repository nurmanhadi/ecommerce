package unit

import (
	"ecommerce/config"
	"ecommerce/internal/repository"
	"ecommerce/internal/service"
	"ecommerce/pkg/infrastructure/database/mariadb"
	"ecommerce/pkg/infrastructure/dto"
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/midtrans/midtrans-go"
)

func TestAddTransaction(t *testing.T) {
	config.LoadConfig()
	midtrans.ServerKey = config.Viper.Midtrans.Key
	midtrans.Environment = midtrans.Sandbox
	db := mariadb.Connection()
	defer db.Close()
	validator := validator.New()
	pRepo := repository.NewProductRepository(db, ctx)
	oRepo := repository.NewOrderRepository(db, ctx)
	tranServ := service.NewTransactionService(&oRepo, &pRepo, validator)
	req := &dto.TransactionRequestDto{
		OrderId: "1251f913-589a-489d-a3f6-3efd96067de9",
	}

	response, err := tranServ.AddTransaction(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}
func TestCheckTransaction(t *testing.T) {
	config.LoadConfig()
	midtrans.ServerKey = config.Viper.Midtrans.Key
	midtrans.Environment = midtrans.Sandbox
	db := mariadb.Connection()
	defer db.Close()
	validator := validator.New()
	pRepo := repository.NewProductRepository(db, ctx)
	oRepo := repository.NewOrderRepository(db, ctx)
	tranServ := service.NewTransactionService(&oRepo, &pRepo, validator)
	req := &dto.TransactionRequestDto{
		OrderId: "1251f913-589a-489d-a3f6-3efd96067de9",
	}

	response, err := tranServ.CheckTransaction(&req.OrderId)
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}
