package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort string
}

func SetupEnv() (cfg AppConfig, err error) {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		// Only return error if we're in development mode
		if os.Getenv("APP_ENV") == "dev" {
			return AppConfig{}, err
		}
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		return AppConfig{}, errors.New("HTTP_PORT environment variable not set")
	}

	return AppConfig{ServerPort: httpPort}, nil
}