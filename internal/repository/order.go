package repository

import (
	"context"
	"database/sql"
	"ecommerce/internal/model"
	"ecommerce/pkg/infrastructure/query"
	"fmt"
)

type orderImpl struct {
	db  *sql.DB
	ctx context.Context
}

func NewOrderRepository(db *sql.DB, ctx context.Context) OrderRepository {
	return &orderImpl{db: db, ctx: ctx}
}
func (r *orderImpl) AddOrder(order *model.Order) error {
	stmt, err := r.db.PrepareContext(r.ctx, query.OrderAdd)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(r.ctx, &order.Id, &order.UserId, &order.ProductId, &order.Quantity, &order.GrossAmount, &order.Status, &order.CreatedAt, &order.UpdatedAt)
	if err != nil {
		return err
	}
	return err
}
func (r *orderImpl) GetOrders(userId *string) ([]model.Order, error) {
	var orders []model.Order
	stmt, err := r.db.PrepareContext(r.ctx, query.OrderGetAll)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(r.ctx, &userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		order := new(model.Order)
		err := rows.Scan(&order.Id, &order.UserId, &order.ProductId, &order.Quantity, &order.GrossAmount, &order.Status, &order.CreatedAt, &order.UpdatedAt)
		if err != nil {
			return nil, err
		}
		orders = append(orders, *order)
	}
	return orders, nil
}
func (r *orderImpl) GetOrderById(orderId *string) (*model.Order, error) {
	order := new(model.Order)
	stmt, err := r.db.PrepareContext(r.ctx, query.OrderGetById)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(r.ctx, &orderId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&order.Id, &order.UserId, &order.ProductId, &order.Quantity, &order.GrossAmount, &order.Status, &order.CreatedAt, &order.UpdatedAt)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("order %s not found", *orderId)
	}
	return order, nil
}
func (r *orderImpl) UpdateOrderStatus(orderId *string, status *string) error {
	stmt, err := r.db.PrepareContext(r.ctx, query.OrderUpdateStatus)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(r.ctx, &status, &orderId)
	if err != nil {
		return err
	}
	return nil
}
func (r *orderImpl) CountOrder(orderId *string) (int, error) {
	var count int
	stmt, err := r.db.PrepareContext(r.ctx, query.OrderCount)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(r.ctx, &orderId)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	if rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return 0, err
		}
	}
	return count, nil
}
