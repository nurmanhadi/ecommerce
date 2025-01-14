package main

import (
	"ecommerce/config"
	"ecommerce/pkg/infrastructure/di"
	_ "ecommerce/pkg/infrastructure/logger"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	config.LoadConfig()
	app := fiber.New(fiber.Config{
		AppName:      config.Viper.App.Name,
		IdleTimeout:  time.Minute * time.Duration(config.Viper.Server.Timeout),
		ReadTimeout:  time.Minute * time.Duration(config.Viper.Server.Timeout),
		WriteTimeout: time.Minute * time.Duration(config.Viper.Server.Timeout),
		Prefork:      config.Viper.Server.Prefork,
	})
	app.Use(logger.New())
	di.DiContainer(app)
	log.Printf("version %s", config.Viper.App.Version)
	app.Listen(fmt.Sprintf(":%d", config.Viper.Server.Port))
}
