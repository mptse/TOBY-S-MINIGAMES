package application

import (
	"context"
	"errors"
	"strings"

	"github.com/google/uuid"
	gamesDomain "github.com/toby-minigames/server/internal/modules/games/domain"
	usersDomain "github.com/toby-minigames/server/internal/modules/users/domain"
	apperrors "github.com/toby-minigames/server/internal/shared/errors"
)

// UserService orquesta los casos de uso de perfil de jugador y subperfiles por minijuego.
type UserService struct {
	playerRepo  usersDomain.PlayerRepository
	profileRepo usersDomain.GameProfileRepository
	gameRepo    gamesDomain.GameRepository
}

// NewUserService crea una nueva instancia de UserService.
func NewUserService(
	playerRepo usersDomain.PlayerRepository,
	profileRepo usersDomain.GameProfileRepository,
	gameRepo gamesDomain.GameRepository,
) *UserService {
	return &UserService{
		playerRepo:  playerRepo,
		profileRepo: profileRepo,
		gameRepo:    gameRepo,
	}
}

// GetProfile obtiene la cuenta del jugador autenticado por su UUID.
func (s *UserService) GetProfile(ctx context.Context, playerID uuid.UUID) (*usersDomain.Player, error) {
	player, err := s.playerRepo.FindByID(ctx, playerID)
	if err != nil {
		if errors.Is(err, usersDomain.ErrPlayerNotFound) {
			return nil, apperrors.NotFound("player not found")
		}
		return nil, apperrors.Internal("failed to fetch player profile")
	}
	return player, nil
}

// UpdateProfile actualiza parcialmente nickname, bio y/o avatar_path del jugador autenticado.
func (s *UserService) UpdateProfile(ctx context.Context, playerID uuid.UUID, nickname, bio, avatarPath *string) (*usersDomain.Player, error) {
	player, err := s.GetProfile(ctx, playerID)
	if err != nil {
		return nil, err
	}

	if nickname != nil {
		trimmed := strings.TrimSpace(*nickname)
		if trimmed == "" {
			return nil, apperrors.Validation("nickname cannot be empty", nil)
		}
		player.Nickname = trimmed
	}

	if bio != nil {
		trimmedBio := strings.TrimSpace(*bio)
		player.Bio = &trimmedBio
	}

	if avatarPath != nil {
		trimmedAvatar := strings.TrimSpace(*avatarPath)
		if trimmedAvatar == "" {
			return nil, apperrors.Validation("avatar_path cannot be empty", nil)
		}
		player.AvatarPath = trimmedAvatar
	}

	if err := s.playerRepo.Update(ctx, player); err != nil {
		return nil, apperrors.Internal("failed to update player profile")
	}

	return player, nil
}

// GetGameProfile obtiene el subperfil de un jugador para un minijuego identificado por su slug.
func (s *UserService) GetGameProfile(ctx context.Context, playerID uuid.UUID, gameSlug string) (*usersDomain.GameProfile, error) {
	game, err := s.gameRepo.FindBySlug(ctx, strings.TrimSpace(gameSlug))
	if err != nil {
		if errors.Is(err, gamesDomain.ErrGameNotFound) {
			return nil, apperrors.NotFound("game not found")
		}
		return nil, apperrors.Internal("failed to resolve game by slug")
	}

	profile, err := s.profileRepo.FindByPlayerAndGame(ctx, playerID, game.ID)
	if err != nil {
		if errors.Is(err, usersDomain.ErrGameProfileNotFound) {
			return nil, apperrors.NotFound("game profile not found for this player")
		}
		return nil, apperrors.Internal("failed to fetch game profile")
	}

	return profile, nil
}
