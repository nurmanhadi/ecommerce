package config

import (
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

type appConfig struct {
	App struct {
		Name    string
		Version string
	}
	Server struct {
		Timeout int
		Port    int
		Prefork bool
	}
	Database struct {
		Mariadb struct {
			Host            string
			Port            int
			User            string
			Password        string
			Dbname          string
			MaxIdleConns    int
			MaxOpenConns    int
			ConnMaxIdleTime int
			ConnMaxLifetime int
		}
	}
	Jwt struct {
		Key string
		Exp int
	}
	Midtrans struct {
		MerchanId string
		ClientId  string
		ServerKey string
	}
}

var Viper = new(appConfig)

func LoadConfig() {
	_, filename, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(filename)
	filepath := filepath.Join(basepath, ".")

	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("yaml")
	config.AddConfigPath(filepath)

	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := config.Unmarshal(&Viper); err != nil {
		panic(err)
	}
}
