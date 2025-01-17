package exception

import "errors"

var (
	OrderNotFound              = errors.New("order not found")
	OrderStatusQueryIsRequired = errors.New("query status is required, min 1, max = 20 length")
)
