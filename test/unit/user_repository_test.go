package unit

import (
	"context"
	"ecommerce/config"
	"ecommerce/internal/model"
	"ecommerce/internal/repository"
	"ecommerce/pkg/infrastructure/database/mariadb"
	"fmt"
	"testing"
	"time"
)

var ctx = context.Background()

func TestRegisterUser(t *testing.T) {
	config.LoadConfig()
	db := mariadb.Connection()
	defer db.Close()

	repo := repository.NewUserRepository(db, ctx)
	user := &model.User{
		Id:        "1",
		Name:      "test",
		Email:     "test@test.com",
		Password:  "test",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := repo.RegisterUser(user)
	if err != nil {
		panic(err.Error())
	}
}
func TestCountUser(t *testing.T) {
	config.LoadConfig()
	db := mariadb.Connection()
	defer db.Close()

	repo := repository.NewUserRepository(db, ctx)
	email := "test@test.com"
	count, err := repo.CountUser(&email)
	if err != nil {
		panic(err.Error())
	}
	println(count)
}
func TestGetUserByEmail(t *testing.T) {
	config.LoadConfig()
	db := mariadb.Connection()
	defer db.Close()

	repo := repository.NewUserRepository(db, ctx)
	email := "est@test.com"
	user, err := repo.GetUserByEmail(&email)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(user)
}
func TestGetUserById(t *testing.T) {
	config.LoadConfig()
	db := mariadb.Connection()
	defer db.Close()

	repo := repository.NewUserRepository(db, ctx)
	userId := "1"
	user, err := repo.GetUserById(&userId)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(user)
}
