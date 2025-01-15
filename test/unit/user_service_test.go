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

var validattion = validator.New()

func TestServiceRegisterUser(t *testing.T) {
	config.LoadConfig()
	db := mariadb.Connection()
	defer db.Close()

	repo := repository.NewUserRepository(db, ctx)
	service := service.NewUserService(&repo, validattion)
	req := &dto.UserRegisterRequestDto{
		Name:     "",
		Email:    "",
		Password: "test",
	}
	err := service.UserRegister(req)
	if err != nil {
		panic(err)
	}
}
func TestServiceLoginUser(t *testing.T) {
	config.LoadConfig()
	db := mariadb.Connection()
	defer db.Close()

	repo := repository.NewUserRepository(db, ctx)
	service := service.NewUserService(&repo, validattion)
	req := &dto.UserLoginRequestDto{
		Email:    "test1@test.com",
		Password: "test",
	}
	token, err := service.UserLogin(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(token)
}
