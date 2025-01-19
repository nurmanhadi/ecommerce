package exception

import "errors"

var (
	PaymentNotFound = errors.New("payment not found")
)
