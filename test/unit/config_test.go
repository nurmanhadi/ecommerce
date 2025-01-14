package unit

import (
	"ecommerce/config"
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	config.LoadConfig()
	fmt.Println(config.Viper.Database.Mariadb.Dbname)
}
