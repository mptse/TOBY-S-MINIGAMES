package infrastructure

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/toby-minigames/server/internal/modules/users/domain"
	"gorm.io/gorm"
)

// PostgresPlayerRepository implementa domain.PlayerRepository utilizando GORM y PostgreSQL.
type PostgresPlayerRepository struct {
	db *gorm.DB
}

// NewPostgresPlayerRepository crea una nueva instancia del repositorio de jugadores.
func NewPostgresPlayerRepository(db *gorm.DB) *PostgresPlayerRepository {
	return &PostgresPlayerRepository{db: db}
}

// Create inserta un nuevo jugador en la base de datos.
func (r *PostgresPlayerRepository) Create(ctx context.Context, player *domain.Player) error {
	model := FromDomainPlayer(player)
	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		return fmt.Errorf("failed to create player: %w", err)
	}
	return nil
}

// FindByID busca un jugador por su identificador UUID.
func (r *PostgresPlayerRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Player, error) {
	var model PlayerModel
	if err := r.db.WithContext(ctx).Where("player_id = ?", id).First(&model).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("player not found: %w", err)
		}
		return nil, fmt.Errorf("failed to find player by ID: %w", err)
	}
	return model.ToDomain(), nil
}

// FindByEmail busca un jugador por su correo electronico.
func (r *PostgresPlayerRepository) FindByEmail(ctx context.Context, email string) (*domain.Player, error) {
	var model PlayerModel
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&model).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("player not found by email: %w", err)
		}
		return nil, fmt.Errorf("failed to find player by email: %w", err)
	}
	return model.ToDomain(), nil
}

// Update persiste los cambios de un jugador existente.
func (r *PostgresPlayerRepository) Update(ctx context.Context, player *domain.Player) error {
	model := FromDomainPlayer(player)
	if err := r.db.WithContext(ctx).Save(model).Error; err != nil {
		return fmt.Errorf("failed to update player: %w", err)
	}
	return nil
}
