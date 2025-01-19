package repository

import (
	"context"
	"database/sql"
	"ecommerce/internal/model"
	"ecommerce/pkg/infrastructure/query"
	"fmt"
	"time"
)

type paymentImpl struct {
	db  *sql.DB
	ctx context.Context
}

func NewPaymentRepository(db *sql.DB, ctx context.Context) PaymentRepository {
	return &paymentImpl{db: db, ctx: ctx}
}

func (r *paymentImpl) AddPayment(payment *model.Payment) error {
	stmt, err := r.db.PrepareContext(r.ctx, query.PaymentAdd)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(r.ctx, &payment.Id, &payment.OrderId, &payment.TransactionId, &payment.MerchantId, &payment.TransactionStatus, &payment.SignatureKey, &payment.PaymentType, &payment.GrossAmount, &payment.FraudStatus, &payment.Currency, &payment.TransactionTime)
	if err != nil {
		return err
	}
	return nil
}
func (r *paymentImpl) UpdateStatus(settlementTime *time.Time, status *string, orderId *string) error {
	stmt, err := r.db.PrepareContext(r.ctx, query.PaymentUpdateByOrderId)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(r.ctx, &settlementTime, &status, &orderId)
	if err != nil {
		return err
	}
	return nil
}
func (r *paymentImpl) CountPayment(orderId *string) (int, error) {
	var count int
	stmt, err := r.db.PrepareContext(r.ctx, query.PaymentCountByOrderId)
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
func (r *paymentImpl) GetPaymentByOrderId(orderId *string) (*model.Payment, error) {
	payment := new(model.Payment)
	stmt, err := r.db.PrepareContext(r.ctx, query.PaymentGetByOrderId)
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
		if err := rows.Scan(&payment.Id, &payment.OrderId, &payment.TransactionId, &payment.MerchantId, &payment.TransactionStatus, &payment.SignatureKey, &payment.PaymentType, &payment.GrossAmount, &payment.FraudStatus, &payment.Currency, &payment.SettlementTime, &payment.TransactionTime); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("payment %s not found", *orderId)
	}
	return payment, nil
}
