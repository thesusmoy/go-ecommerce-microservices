package main

import (
	"go-ecommerce-microservices/config"
	"go-ecommerce-microservices/internal/api"
	"log"
)
func main() {
	cfg, err := config.SetupEnv()

	if err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}

	api.StartServer(cfg)
}
