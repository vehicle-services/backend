package config

import (
	"github.com/joho/godotenv"
	"os"
)

var Envs = initConfig()

type Config struct {
	Port       string
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
}

func initConfig() Config {
	godotenv.Load()
	return Config{
		Port:       getEnv("PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "Shivam29@"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBName:     getEnv("DB_NAME", "mechanix"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
