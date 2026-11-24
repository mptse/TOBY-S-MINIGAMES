package config

import (
	"os"
	"testing"
)

func TestLoad_Defaults(t *testing.T) {
	os.Clearenv()

	cfg := Load()

	if cfg.Port != "8080" {
		t.Errorf("expected Port=8080, got %s", cfg.Port)
	}
	if cfg.Environment != "development" {
		t.Errorf("expected Environment=development, got %s", cfg.Environment)
	}
	if cfg.DatabaseURL != "" {
		t.Errorf("expected empty default DatabaseURL, got %s", cfg.DatabaseURL)
	}
}

func TestLoad_CustomEnv(t *testing.T) {
	os.Setenv("PORT", "9090")
	os.Setenv("ENVIRONMENT", "production")
	os.Setenv("DATABASE_URL", "postgres://test:postgres@localhost:5432/testdb")
	defer os.Clearenv()

	cfg := Load()

	if cfg.Port != "9090" {
		t.Errorf("expected Port=9090, got %s", cfg.Port)
	}
	if cfg.Environment != "production" {
		t.Errorf("expected Environment=production, got %s", cfg.Environment)
	}
	if cfg.DatabaseURL != "postgres://test:postgres@localhost:5432/testdb" {
		t.Errorf("expected custom DatabaseURL, got %s", cfg.DatabaseURL)
	}
}
