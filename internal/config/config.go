package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DatabaseDSN string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Not found file .env")
	}

	return &Config{
		Port: getEnv("PORT", "8080"),
		DatabaseDSN: "host=" + getEnv("DB_HOST", "localhost") +
			" user=" + getEnv("DB_USER", "postgres") +
			" password=" + getEnv("DB_PASSWORD", "") +
			" dbname=" + getEnv("DB_NAME", "postgres") +
			" port=" + getEnv("DB_PORT", "5432") +
			" sslmode=" + getEnv("DB_SSLMODE", "disable"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
