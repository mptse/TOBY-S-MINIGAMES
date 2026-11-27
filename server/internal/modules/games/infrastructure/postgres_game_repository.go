package infrastructure

import (
	"context"
	"errors"
	"fmt"

	"github.com/toby-minigames/server/internal/modules/games/domain"
	"gorm.io/gorm"
)

// PostgresGameRepository implementa domain.GameRepository utilizando GORM y PostgreSQL.
type PostgresGameRepository struct {
	db *gorm.DB
}

// NewPostgresGameRepository crea una nueva instancia del repositorio de minijuegos.
func NewPostgresGameRepository(db *gorm.DB) *PostgresGameRepository {
	return &PostgresGameRepository{db: db}
}

// Create inserta un nuevo minijuego en el catalogo y actualiza su ID autogenerado.
func (r *PostgresGameRepository) Create(ctx context.Context, game *domain.Game) error {
	model := FromDomainGame(game)
	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		return fmt.Errorf("failed to create game: %w", err)
	}
	game.ID = model.GameID
	return nil
}

// FindByID busca un minijuego por su identificador numerico junto con sus categorias.
func (r *PostgresGameRepository) FindByID(ctx context.Context, id int) (*domain.Game, error) {
	var model GameModel
	if err := r.db.WithContext(ctx).Preload("Categories").Where("game_id = ?", id).First(&model).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("%w: %v", domain.ErrGameNotFound, err)
		}
		return nil, fmt.Errorf("failed to find game by ID: %w", err)
	}
	return model.ToDomain(), nil
}

// FindBySlug busca un minijuego por su slug unico junto con sus categorias.
func (r *PostgresGameRepository) FindBySlug(ctx context.Context, slug string) (*domain.Game, error) {
	var model GameModel
	if err := r.db.WithContext(ctx).Preload("Categories").Where("name_slug = ?", slug).First(&model).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("%w: %v", domain.ErrGameNotFound, err)
		}
		return nil, fmt.Errorf("failed to find game by slug: %w", err)
	}
	return model.ToDomain(), nil
}

// ListActive retorna todos los minijuegos activos del catalogo.
func (r *PostgresGameRepository) ListActive(ctx context.Context) ([]*domain.Game, error) {
	var models []GameModel
	if err := r.db.WithContext(ctx).Preload("Categories").Where("is_active = ?", true).Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to list active games: %w", err)
	}

	games := make([]*domain.Game, len(models))
	for i := range models {
		games[i] = models[i].ToDomain()
	}
	return games, nil
}

// ListActiveByCategory retorna los minijuegos activos filtrados por el slug de una categoria.
func (r *PostgresGameRepository) ListActiveByCategory(ctx context.Context, categorySlug string) ([]*domain.Game, error) {
	var models []GameModel
	err := r.db.WithContext(ctx).
		Preload("Categories").
		Joins("JOIN game_categories gc ON gc.game_id = games.game_id").
		Joins("JOIN categories c ON c.category_id = gc.category_id").
		Where("games.is_active = ? AND c.category_slug = ?", true, categorySlug).
		Find(&models).Error
	if err != nil {
		return nil, fmt.Errorf("failed to list active games by category: %w", err)
	}

	games := make([]*domain.Game, len(models))
	for i := range models {
		games[i] = models[i].ToDomain()
	}
	return games, nil
}
