package config

import (
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port       string
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
}

func LoadConfig() AppConfig {
	if err := godotenv.Load("./env/.env"); err != nil {
		return AppConfig{}
	}
	return AppConfig{
		Port:       os.Getenv("PORT"),
		DBHost:     os.Getenv("DBHost"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
	}
}
