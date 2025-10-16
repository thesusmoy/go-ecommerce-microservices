package api

import (
	"go-ecommerce-microservices/config"

	"github.com/gofiber/fiber/v2"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	app.Listen(config.ServerPort)
}

