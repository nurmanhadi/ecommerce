package di

import (
	"context"
	"ecommerce/config"
	"ecommerce/internal/controller"
	"ecommerce/internal/repository"
	"ecommerce/internal/routes"
	"ecommerce/internal/service"
	"ecommerce/pkg/infrastructure/database/mariadb"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/midtrans/midtrans-go"
)

func DiContainer(app *fiber.App) {
	midtrans.ServerKey = config.Viper.Midtrans.Key
	midtrans.Environment = midtrans.Sandbox
	ctx := context.Background()
	validation := validator.New()
	db := mariadb.Connection()

	userRepo := repository.NewUserRepository(db, ctx)
	userServ := service.NewUserService(&userRepo, validation)
	userCont := controller.NewUserController(&userServ)
	routes.UserRoute(app, userCont)

	productRepo := repository.NewProductRepository(db, ctx)
	productServ := service.NewProductService(&productRepo, validation)
	productCont := controller.NewProductController(&productServ)
	routes.ProductRoute(app, productCont)

	orderRepo := repository.NewOrderRepository(db, ctx)
	orderServ := service.NewOrderService(&orderRepo, &productRepo, &userRepo, validation)
	orderCont := controller.NewOrderController(&orderServ)
	routes.OrderRoute(app, orderCont)

	transactionServ := service.NewTransactionService(&orderRepo, &productRepo, validation)
	transactionCont := controller.NewTransactionController(&transactionServ)
	routes.TransactionRoute(app, transactionCont)

	payRepo := repository.NewPaymentRepository(db, ctx)
	payServ := service.NewPaymentService(&payRepo, &productRepo, &orderRepo, validation)
	payCont := controller.NewPaymentController(&payServ)
	routes.PaymentRoute(app, payCont)

	notifCont := controller.NewNotificationCOntroller(&payServ)
	routes.NotificationRoute(app, notifCont)
}
