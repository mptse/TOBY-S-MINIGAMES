package main

import (
	"context"
	"log"
	"time"

	"github.com/toby-minigames/server/internal/infrastructure/config"
	httpserver "github.com/toby-minigames/server/internal/infrastructure/http"
	"github.com/toby-minigames/server/internal/shared/database"
	"github.com/toby-minigames/server/internal/shared/security"
)

func main() {
	cfg := config.Load()

	log.Printf("Starting Toby's Minigames Backend in environment: %s (port: %s)", cfg.Environment, cfg.Port)

	if cfg.DatabaseURL == "" {
		log.Fatal("[FATAL] DATABASE_URL must be configured")
	}
	if cfg.JWTSecret == "" {
		log.Fatal("[FATAL] JWT_SECRET must be configured")
	}

	db, err := database.NewConnection(cfg.DatabaseURL, cfg.Environment, database.DefaultPoolConfig())
	if err != nil {
		log.Fatalf("[FATAL] Could not connect to database: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := database.Ping(ctx, db); err != nil {
		log.Fatalf("[FATAL] Database ping failed: %v", err)
	}

	jwtMgr, err := security.NewJWTManager(cfg.JWTSecret, cfg.JWTExpirationHours)
	if err != nil {
		log.Fatalf("[FATAL] Failed to initialize JWT manager: %v", err)
	}

	router := httpserver.NewRouter(db, jwtMgr, cfg.Environment)
	log.Printf("[OK] HTTP server listening on :%s", cfg.Port)

	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("[FATAL] Server terminated with error: %v", err)
	}
}
