package unit

import (
	"ecommerce/config"
	"ecommerce/pkg/infrastructure/database/mariadb"
	"testing"
)

func TestXxx(t *testing.T) {
	config.LoadConfig()
	db := mariadb.Connection()
	defer db.Close()
}
