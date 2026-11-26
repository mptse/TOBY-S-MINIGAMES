package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config estructura las variables de configuracion de la aplicacion backend.
type Config struct {
	Port               string
	Environment        string
	DatabaseURL        string
	JWTSecret          string
	JWTExpirationHours int
}

// Load lee las variables de entorno desde el sistema o archivo .env si existe.
func Load() *Config {
	_ = godotenv.Load()

	expHours := 24
	if raw := getEnv("JWT_EXPIRATION_HOURS", ""); raw != "" {
		if parsed, err := strconv.Atoi(raw); err == nil && parsed > 0 {
			expHours = parsed
		}
	}

	return &Config{
		Port:               getEnv("PORT", "8080"),
		Environment:        getEnv("ENVIRONMENT", "development"),
		DatabaseURL:        getEnv("DATABASE_URL", ""),
		JWTSecret:          getEnv("JWT_SECRET", ""),
		JWTExpirationHours: expHours,
	}
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok && val != "" {
		return val
	}
	return fallback
}
