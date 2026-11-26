package main

import (
	"context"
	"log"
	"time"

	"github.com/toby-minigames/server/internal/infrastructure/config"
	"github.com/toby-minigames/server/internal/shared/database"
)

func main() {
	cfg := config.Load()

	log.Printf("Starting Toby's Minigames Backend in environment: %s (port: %s)", cfg.Environment, cfg.Port)

	if cfg.DatabaseURL == "" {
		log.Println("[INFO] DATABASE_URL not configured, skipping initial DB connection.")
		return
	}

	db, err := database.NewConnection(cfg.DatabaseURL, cfg.Environment, database.DefaultPoolConfig())
	if err != nil {
		log.Printf("[WARN] Could not connect to database on startup: %v", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := database.Ping(ctx, db); err != nil {
		log.Printf("[WARN] Database ping failed: %v", err)
		return
	}

	log.Println("[OK] PostgreSQL connection and ping verified successfully.")
}
