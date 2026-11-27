package infrastructure

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/toby-minigames/server/internal/modules/users/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// PostgresGameProfileRepository implementa domain.GameProfileRepository utilizando GORM y PostgreSQL.
type PostgresGameProfileRepository struct {
	db *gorm.DB
}

// NewPostgresGameProfileRepository crea una nueva instancia del repositorio de subperfiles por juego.
func NewPostgresGameProfileRepository(db *gorm.DB) *PostgresGameProfileRepository {
	return &PostgresGameProfileRepository{db: db}
}

// Upsert inserta o actualiza el subperfil de un jugador para un minijuego especifico.
func (r *PostgresGameProfileRepository) Upsert(ctx context.Context, profile *domain.GameProfile) error {
	model, err := FromDomainGameProfile(profile)
	if err != nil {
		return fmt.Errorf("failed to map game profile to model: %w", err)
	}

	err = r.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "player_id"}, {Name: "game_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"highest_score", "matches_played", "game_data", "updated_at"}),
	}).Create(model).Error
	if err != nil {
		return fmt.Errorf("failed to upsert game profile: %w", err)
	}

	return nil
}

// FindByPlayerAndGame obtiene el subperfil de un jugador para un minijuego dado.
func (r *PostgresGameProfileRepository) FindByPlayerAndGame(ctx context.Context, playerID uuid.UUID, gameID int) (*domain.GameProfile, error) {
	var model GameProfileModel
	err := r.db.WithContext(ctx).
		Where("player_id = ? AND game_id = ?", playerID, gameID).
		First(&model).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("%w: %v", domain.ErrGameProfileNotFound, err)
		}
		return nil, fmt.Errorf("failed to find game profile: %w", err)
	}

	return model.ToDomain()
}

// ListByPlayer lista todos los subperfiles asociados a un jugador.
func (r *PostgresGameProfileRepository) ListByPlayer(ctx context.Context, playerID uuid.UUID) ([]*domain.GameProfile, error) {
	var models []GameProfileModel
	if err := r.db.WithContext(ctx).Where("player_id = ?", playerID).Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to list game profiles by player: %w", err)
	}

	profiles := make([]*domain.GameProfile, 0, len(models))
	for i := range models {
		profile, err := models[i].ToDomain()
		if err != nil {
			return nil, fmt.Errorf("failed to map game profile model: %w", err)
		}
		profiles = append(profiles, profile)
	}

	return profiles, nil
}
