package query

var (
	UserLogin      = "INSERT users(id, name, email, password, created_at, updated_at) VALUES(?,?,?,?,?,?)"
	UserCount      = "SELECT COUNT(*) FROM users WHERE email = ?"
	UserCountById  = "SELECT COUNT(*) FROM users WHERE id = ?"
	UserGetByEmail = "SELECT id, name, email, password, created_at, updated_at FROM users WHERE email = ?"
	UserGetById    = "SELECT id, name, email, password, created_at, updated_at FROM users WHERE id = ?"
)
