package unit

import (
	"ecommerce/config"
	"ecommerce/internal/model"
	"ecommerce/internal/repository"
	"ecommerce/pkg/infrastructure/database/mariadb"
	"testing"
	"time"
)

func TestPaymentAdd(t *testing.T) {
	config.LoadConfig()
	db := mariadb.Connection()
	defer db.Close()
	repo := repository.NewPaymentRepository(db, ctx)
	payment := &model.Payment{
		Id:                "1",
		OrderId:           "1",
		TransactionId:     "1",
		MerchantId:        "1",
		TransactionStatus: "pending",
		SignatureKey:      "1",
		PaymentType:       "bank",
		GrossAmount:       "30000",
		FraudStatus:       "accept",
		Currency:          "IDR",
		SettlementTime:    time.Now(),
		TransactionTime:   time.Now(),
	}
	err := repo.AddPayment(payment)
	if err != nil {
		panic(err)
	}
}
