package repository

import (
	"context"
	"database/sql"
	"ecommerce/internal/model"
	"ecommerce/pkg/infrastructure/dto"
	"ecommerce/pkg/infrastructure/query"
	"fmt"
	"time"
)

type productImpl struct {
	db  *sql.DB
	ctx context.Context
}

func NewProductRepository(db *sql.DB, ctx context.Context) ProductRepository {
	return &productImpl{db: db, ctx: ctx}
}
func (r *productImpl) AddProduct(product *model.Product) error {
	stmt, err := r.db.PrepareContext(r.ctx, query.ProductAdd)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(r.ctx, &product.Name, &product.Description, &product.Price, &product.Stock, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}
func (r *productImpl) GetProductById(productId *int) (*model.Product, error) {
	product := new(model.Product)
	stmt, err := r.db.PrepareContext(r.ctx, query.ProductGetById)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(r.ctx, &productId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Stock, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("product %d not found", product.Id)
	}
	return product, nil
}
func (r *productImpl) UpdateProduct(productId *int, product *dto.ProductUpdateRequestDto) error {
	value := "UPDATE products SET "
	args := []interface{}{}
	if product.Name != nil {
		value += "name = ?, "
		args = append(args, product.Name)
	}
	if product.Description != nil {
		value += "description = ?, "
		args = append(args, product.Description)
	}
	if product.Price != nil {
		value += "price = ?, "
		args = append(args, product.Price)
	}
	if product.Stock != nil {
		value += "stock = ?, "
		args = append(args, product.Stock)
	}
	value += "updated_at = ?, "
	args = append(args, time.Now())
	value = value[:len(value)-2] + " WHERE id = ?"
	args = append(args, productId)

	stmt, err := r.db.PrepareContext(r.ctx, value)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(r.ctx, args...)
	if err != nil {
		return err
	}
	return nil
}
func (r *productImpl) DeleteProduct(productId *int) error {
	stmt, err := r.db.PrepareContext(r.ctx, query.ProductDelete)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(r.ctx, &productId)
	if err != nil {
		return err
	}
	return nil
}
func (r *productImpl) GetProducts() ([]model.Product, error) {
	var products []model.Product
	stmt, err := r.db.PrepareContext(r.ctx, query.ProductGetAll)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(r.ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		product := new(model.Product)
		err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Stock, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, *product)
	}
	return products, nil
}
func (r *productImpl) CountProduct(productId *int) (int, error) {
	var count int
	stmt, err := r.db.PrepareContext(r.ctx, query.ProductCount)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(r.ctx, &productId)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			return 0, err
		}
	} else {
		return 0, fmt.Errorf("error query sql")
	}
	return count, nil
}
