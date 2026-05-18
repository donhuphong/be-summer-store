package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DatabaseDSN string
	R2CloudFare r2CloudFareCfg
	JWT         jWT
}

type r2CloudFareCfg struct {
	R2AccessKeyID     string
	R2SecretAccessKey string
	R2BucketName      string
	R2PublicURL       string
	R2AccountID       string
}

type jWT struct {
	AccessSecret  string
	RefreshSecret string
	PASS          string
}

var AppConfig Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Not found file .env")
	}

	AppConfig = Config{
		Port:        getEnv("PORT", "8080"),
		DatabaseDSN: buildDSN(),

		R2CloudFare: r2CloudFareCfg{
			R2AccessKeyID:     getEnv("R2_ACCESS_KEY_ID", ""),
			R2SecretAccessKey: getEnv("R2_SECRET_ACCESS_KEY", ""),
			R2BucketName:      getEnv("R2_BUCKET_NAME", ""),
			R2PublicURL:       getEnv("R2_PUBLIC_URL", ""),
			R2AccountID:       getEnv("R2_ACCOUNT_ID", ""),
		},

		JWT: jWT{
			AccessSecret:  getEnv("JWT_ACCESS_SECRET", ""),
			RefreshSecret: getEnv("JWT_REFRESH_SECRET", ""),
			PASS:          getEnv("PASS", ""),
		},
	}
	validateConfig()
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func validateConfig() {
	if AppConfig.JWT.AccessSecret == "" || AppConfig.JWT.RefreshSecret == "" || AppConfig.JWT.PASS == "" {
		log.Fatal("AccessSecret and RefreshSecret is required")
	}
}

func buildDSN() string {
	return "host=" + getEnv("DB_HOST", "localhost") +
		" user=" + getEnv("DB_USER", "postgres") +
		" password=" + getEnv("DB_PASSWORD", "") +
		" dbname=" + getEnv("DB_NAME", "postgres") +
		" port=" + getEnv("DB_PORT", "5432") +
		" sslmode=" + getEnv("DB_SSLMODE", "disable")
}
