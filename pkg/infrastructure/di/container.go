package di

import (
	"context"
	"ecommerce/internal/controller"
	"ecommerce/internal/repository"
	"ecommerce/internal/routes"
	"ecommerce/internal/service"
	"ecommerce/pkg/infrastructure/database/mariadb"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func DiContainer(app *fiber.App) {
	ctx := context.Background()
	validation := validator.New()
	db := mariadb.Connection()

	userRepo := repository.NewUserRepository(db, ctx)
	userServ := service.NewUserService(&userRepo, validation)
	userCont := controller.NewUserController(&userServ)
	routes.UserRoute(app, userCont)
}
