package application

import (
	"context"
	"errors"
	"strings"

	gamesDomain "github.com/toby-minigames/server/internal/modules/games/domain"
	appErrors "github.com/toby-minigames/server/internal/shared/errors"
)

// GameService orquesta las consultas del catalogo de minijuegos y categorias.
type GameService struct {
	gameRepo     gamesDomain.GameRepository
	categoryRepo gamesDomain.CategoryRepository
}

// NewGameService crea una nueva instancia de GameService.
func NewGameService(gameRepo gamesDomain.GameRepository, categoryRepo gamesDomain.CategoryRepository) *GameService {
	return &GameService{
		gameRepo:     gameRepo,
		categoryRepo: categoryRepo,
	}
}

// ListGames retorna los minijuegos activos, filtrados opcionalmente por slug de categoria.
func (s *GameService) ListGames(ctx context.Context, categorySlug string) ([]*gamesDomain.Game, error) {
	categorySlug = strings.TrimSpace(categorySlug)
	if categorySlug != "" {
		if _, err := s.categoryRepo.FindBySlug(ctx, categorySlug); err != nil {
			if errors.Is(err, gamesDomain.ErrCategoryNotFound) {
				return nil, appErrors.NotFound("category not found")
			}
			return nil, appErrors.Internal("failed to verify category")
		}

		games, err := s.gameRepo.ListActiveByCategory(ctx, categorySlug)
		if err != nil {
			return nil, appErrors.Internal("failed to list games by category")
		}
		return games, nil
	}

	games, err := s.gameRepo.ListActive(ctx)
	if err != nil {
		return nil, appErrors.Internal("failed to list active games")
	}
	return games, nil
}

// GetGameBySlug retorna el detalle completo de un minijuego y sus categorias.
func (s *GameService) GetGameBySlug(ctx context.Context, slug string) (*gamesDomain.Game, error) {
	game, err := s.gameRepo.FindBySlug(ctx, strings.TrimSpace(slug))
	if err != nil {
		if errors.Is(err, gamesDomain.ErrGameNotFound) {
			return nil, appErrors.NotFound("game not found")
		}
		return nil, appErrors.Internal("failed to fetch game detail")
	}
	return game, nil
}

// ListCategories retorna todas las categorias disponibles ordenadas alfabeticamente.
func (s *GameService) ListCategories(ctx context.Context) ([]*gamesDomain.Category, error) {
	categories, err := s.categoryRepo.ListAll(ctx)
	if err != nil {
		return nil, appErrors.Internal("failed to list categories")
	}
	return categories, nil
}
