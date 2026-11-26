package database

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// PoolConfig define los parametros del pool de conexiones SQL.
type PoolConfig struct {
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

// DefaultPoolConfig retorna la configuracion recomendada para el pool de conexiones.
func DefaultPoolConfig() PoolConfig {
	return PoolConfig{
		MaxIdleConns:    10,
		MaxOpenConns:    100,
		ConnMaxLifetime: time.Hour,
	}
}

// NewConnection abre la conexion a PostgreSQL a traves de GORM utilizando DATABASE_URL.
func NewConnection(databaseURL, environment string, poolCfg PoolConfig) (*gorm.DB, error) {
	if databaseURL == "" {
		return nil, fmt.Errorf("databaseURL cannot be empty")
	}

	logLevel := logger.Warn
	if environment == "development" {
		logLevel = logger.Info
	}

	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	sqlDB.SetMaxIdleConns(poolCfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(poolCfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(poolCfg.ConnMaxLifetime)

	return db, nil
}

// Ping verifica la conectividad activa con la base de datos PostgreSQL.
func Ping(ctx context.Context, db *gorm.DB) error {
	if db == nil {
		return fmt.Errorf("database handle cannot be nil")
	}

	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying sql.DB for ping: %w", err)
	}

	if err := sqlDB.PingContext(ctx); err != nil {
		return fmt.Errorf("database ping failed: %w", err)
	}

	return nil
}
