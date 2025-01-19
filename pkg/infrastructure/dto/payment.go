package dto

type PaymentRequestDto struct {
	TransactionTime   string `json:"transaction_time" validate:"required"`
	TransactionStatus string `json:"transaction_status" validate:"required"`
	TransactionId     string `json:"transaction_id" validate:"required"`
	StatusMessage     string `json:"status_message" validate:"required"`
	StatusCOde        string `json:"status_code" validate:"required"`
	SignatureKey      string `json:"signature_key" validate:"required"`
	SettlementTime    string `json:"settlement_time"`
	PaymentType       string `json:"payment_type" validate:"required"`
	OrderId           string `json:"order_id" validate:"required"`
	MerchantId        string `json:"merchant_id" validate:"required"`
	GrossAmount       string `json:"gross_amount" validate:"required"`
	FraudStatus       string `json:"fraud_status" validate:"required"`
	Currency          string `json:"currency" validate:"required"`
}
