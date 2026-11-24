package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config estructura las variables de configuracion de la aplicacion backend.
type Config struct {
	Port        string
	Environment string
	DatabaseURL string
	JWTSecret   string
}

// Load lee las variables de entorno desde el sistema o archivo .env si existe.
func Load() *Config {
	_ = godotenv.Load()

	return &Config{
		Port:        getEnv("PORT", "8080"),
		Environment: getEnv("ENVIRONMENT", "development"),
		DatabaseURL: getEnv("DATABASE_URL", ""),
		JWTSecret:   getEnv("JWT_SECRET", ""),
	}
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok && val != "" {
		return val
	}
	return fallback
}
