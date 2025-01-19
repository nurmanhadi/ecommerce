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
)

func TestOrderAdd(t *testing.T) {
	config.LoadConfig()
	db := mariadb.Connection()
	defer db.Close()
	validator := validator.New()
	pRepo := repository.NewProductRepository(db, ctx)
	uRepo := repository.NewUserRepository(db, ctx)
	oRepo := repository.NewOrderRepository(db, ctx)

	serv := service.NewOrderService(&oRepo, &pRepo, &uRepo, validator)
	userId := "1"
	order := &dto.OrderAddRequestDto{
		ProductId: 6,
		Quantity:  5,
	}
	id, err := serv.AddOrder(&userId, order)
	if err != nil {
		panic(err)
	}
	fmt.Println(id)
}
func TestOrderGetAll(t *testing.T) {
	config.LoadConfig()
	db := mariadb.Connection()
	defer db.Close()
	validator := validator.New()
	pRepo := repository.NewProductRepository(db, ctx)
	uRepo := repository.NewUserRepository(db, ctx)
	oRepo := repository.NewOrderRepository(db, ctx)

	serv := service.NewOrderService(&oRepo, &pRepo, &uRepo, validator)
	userId := "1"
	orders, err := serv.GetOrders(&userId)
	if err != nil {
		panic(err)
	}
	fmt.Println(orders)
}
func TestOrderGetById(t *testing.T) {
	config.LoadConfig()
	db := mariadb.Connection()
	defer db.Close()
	validator := validator.New()
	pRepo := repository.NewProductRepository(db, ctx)
	uRepo := repository.NewUserRepository(db, ctx)
	oRepo := repository.NewOrderRepository(db, ctx)

	serv := service.NewOrderService(&oRepo, &pRepo, &uRepo, validator)
	orderId := "bd620557-080b-407e-b8c0-ee950721edad"
	order, err := serv.GetOrderById(&orderId)
	if err != nil {
		panic(err)
	}
	fmt.Println(order)
}
func TestOrderUpdateStatus(t *testing.T) {
	config.LoadConfig()
	db := mariadb.Connection()
	defer db.Close()
	validator := validator.New()
	pRepo := repository.NewProductRepository(db, ctx)
	uRepo := repository.NewUserRepository(db, ctx)
	oRepo := repository.NewOrderRepository(db, ctx)

	serv := service.NewOrderService(&oRepo, &pRepo, &uRepo, validator)
	orderId := "1251f913-589a-489d-a3f6-3efd96067de9"
	status := "pending"
	err := serv.UpdateOrderStatus(&orderId, &status)
	if err != nil {
		panic(err)
	}
}
