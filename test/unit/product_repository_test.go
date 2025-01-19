package unit

import (
	"ecommerce/config"
	"ecommerce/internal/model"
	"ecommerce/internal/repository"
	"ecommerce/pkg/infrastructure/database/mariadb"
	"ecommerce/pkg/infrastructure/dto"
	"fmt"
	"testing"
	"time"
)

func TestProductAdd(t *testing.T) {
	config.LoadConfig()
	db := mariadb.Connection()
	defer db.Close()
	repo := repository.NewProductRepository(db, ctx)
	product := model.Product{
		Name:        "hura",
		Description: "hehe",
		Price:       1000,
		Stock:       100,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	id, err := repo.AddProduct(&product)
	if err != nil {
		panic(err)
	}
	fmt.Println(id)
}
func TestProductUpdate(t *testing.T) {
	config.LoadConfig()
	db := mariadb.Connection()
	defer db.Close()
	repo := repository.NewProductRepository(db, ctx)
	st := 75000000
	product := &dto.ProductUpdateRequestDto{
		Stock: &st,
	}
	id := 2
	err := repo.UpdateProduct(&id, product)
	if err != nil {
		panic(err)
	}
}
func TestProductDelete(t *testing.T) {
	config.LoadConfig()
	db := mariadb.Connection()
	defer db.Close()
	repo := repository.NewProductRepository(db, ctx)

	id := 2
	err := repo.DeleteProduct(&id)
	if err != nil {
		panic(err)
	}
}
