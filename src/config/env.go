package config

import (
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func LoadEnv(key string) (env string) {
	err := godotenv.Load(".env")
	if err != nil {
		LogError("error loading .env file", zap.String("error", err.Error()))
	}

	return os.Getenv(key)
}
