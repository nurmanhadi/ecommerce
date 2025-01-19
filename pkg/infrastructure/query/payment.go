package query

var (
	PaymentAdd             = "INSERT INTO payments(id, order_id, transaction_id, merchant_id, transaction_status, signature_key, payment_type, gross_amount, fraud_status, currency, transaction_time) VALUES(?,?,?,?,?,?,?,?,?,?,?)"
	PaymentCountByOrderId  = "SELECT COUNT(*) FROM payments WHERE order_id = ?"
	PaymentUpdateByOrderId = "UPDATE payments SET  settlement_time = ?, transaction_status = ? WHERE order_id = ?"
	PaymentGetByOrderId    = "SELECT id, order_id, transaction_id, merchant_id, transaction_status, signature_key, payment_type, gross_amount, fraud_status, currency, settlement_time, transaction_time FROM payments WHERE order_id = ?"
)
