package main

import (
	"log"

	"github.com/toby-minigames/server/internal/infrastructure/config"
	"github.com/toby-minigames/server/internal/shared/database"
)

func main() {
	cfg := config.Load()

	log.Printf("Starting Toby's Minigames Backend in environment: %s (port: %s)", cfg.Environment, cfg.Port)

	if cfg.DatabaseURL != "" {
		db, err := database.NewConnection(cfg.DatabaseURL, database.DefaultPoolConfig())
		if err != nil {
			log.Printf("[WARN] Could not connect to database on startup: %v", err)
		} else {
			log.Println("[OK] PostgreSQL connection initialized successfully.")
			_ = db
		}
	} else {
		log.Println("[INFO] DATABASE_URL not configured, skipping initial DB connection.")
	}
}
