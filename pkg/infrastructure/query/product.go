package query

var (
	ProductAdd     = "INSERT products(name, description, price, stock, created_at, updated_at) VALUES(?,?,?,?,?,?)"
	ProductCount   = "SELECT COUNT(*) FROM products WHERE id = ?"
	ProductGetById = "SELECT id, name, description, price, stock, created_at, updated_at FROM products WHERE id = ?"
	ProductGetAll  = "SELECT id, name, description, price, stock, created_at, updated_at FROM products"
	ProductDelete  = "DELETE FROM products WHERE id = ?"
)
