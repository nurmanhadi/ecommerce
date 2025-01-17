package query

var (
	OrderAdd          = "INSERT INTO orders(id, user_id, product_id, quantity, gross_amount, status, created_at, updated_at) VALUES(?,?,?,?,?,?,?,?)"
	OrderGetAll       = "SELECT id, user_id, product_id, quantity, gross_amount, status, created_at, updated_at FROM orders WHERE user_id = ?"
	OrderGetById      = "SELECT id, user_id, product_id, quantity, gross_amount, status, created_at, updated_at FROM orders WHERE id = ?"
	OrderUpdateStatus = "UPDATE orders SET status = ? WHERE id = ?"
	OrderCount        = "SELECT COUNT(*) FROM orders WHERE id = ?"
)
