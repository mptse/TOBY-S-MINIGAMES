package application

import (
	"context"
	"errors"
	"strings"

	gamesDomain "github.com/toby-minigames/server/internal/modules/games/domain"
	apperrors "github.com/toby-minigames/server/internal/shared/errors"
)

// GameService orquesta las consultas del catalogo de minijuegos.
type GameService struct {
	gameRepo gamesDomain.GameRepository
}

// NewGameService crea una nueva instancia de GameService.
func NewGameService(gameRepo gamesDomain.GameRepository) *GameService {
	return &GameService{gameRepo: gameRepo}
}

// ListGames retorna los minijuegos activos, filtrados opcionalmente por slug de categoria.
func (s *GameService) ListGames(ctx context.Context, categorySlug string) ([]*gamesDomain.Game, error) {
	categorySlug = strings.TrimSpace(categorySlug)
	if categorySlug != "" {
		games, err := s.gameRepo.ListActiveByCategory(ctx, categorySlug)
		if err != nil {
			return nil, apperrors.Internal("failed to list games by category")
		}
		return games, nil
	}

	games, err := s.gameRepo.ListActive(ctx)
	if err != nil {
		return nil, apperrors.Internal("failed to list active games")
	}
	return games, nil
}

// GetGameBySlug retorna el detalle completo de un minijuego y sus categorias.
func (s *GameService) GetGameBySlug(ctx context.Context, slug string) (*gamesDomain.Game, error) {
	game, err := s.gameRepo.FindBySlug(ctx, strings.TrimSpace(slug))
	if err != nil {
		if errors.Is(err, gamesDomain.ErrGameNotFound) {
			return nil, apperrors.NotFound("game not found")
		}
		return nil, apperrors.Internal("failed to fetch game detail")
	}
	return game, nil
}
