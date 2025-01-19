package model

import "time"

type Payment struct {
	Id                string    `json:"id"`
	OrderId           string    `json:"order_id"`
	TransactionId     string    `json:"transaction_id"`
	MerchantId        string    `json:"merchant_id"`
	TransactionStatus string    `json:"transaction_status"`
	SignatureKey      string    `json:"signature_key"`
	PaymentType       string    `json:"payment_type"`
	GrossAmount       string    `json:"gross_amount"`
	FraudStatus       string    `json:"fraud_status"`
	Currency          string    `json:"currency"`
	SettlementTime    time.Time `json:"settlement_time"`
	TransactionTime   time.Time `json:"transaction_time"`
}
