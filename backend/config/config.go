package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DBHost string
	DBName string
	DBUser string
	DBPass string
}

func LoadConfig() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		return Config{}, err
	}

	cfg := Config{
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
	}

	return cfg, nil
}
