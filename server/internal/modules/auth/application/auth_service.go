package application

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	usersDomain "github.com/toby-minigames/server/internal/modules/users/domain"
	apperrors "github.com/toby-minigames/server/internal/shared/errors"
	"github.com/toby-minigames/server/internal/shared/security"
)

// AuthService orquesta los casos de uso de registro, login e invitados.
type AuthService struct {
	playerRepo usersDomain.PlayerRepository
	jwtMgr     *security.JWTManager
}

// NewAuthService crea una nueva instancia de AuthService.
func NewAuthService(playerRepo usersDomain.PlayerRepository, jwtMgr *security.JWTManager) *AuthService {
	return &AuthService{
		playerRepo: playerRepo,
		jwtMgr:     jwtMgr,
	}
}

// GuestSession representa los datos efimeros de una sesion de invitado.
type GuestSession struct {
	GuestID  string
	Nickname string
	Role     string
}

// Register registra un nuevo jugador, valida unicidad de correo y retorna JWT + jugador.
func (s *AuthService) Register(ctx context.Context, email, password, nickname string) (string, *usersDomain.Player, error) {
	email = strings.ToLower(strings.TrimSpace(email))
	nickname = strings.TrimSpace(nickname)

	if email == "" || !strings.Contains(email, "@") || len(password) < 6 || nickname == "" {
		return "", nil, apperrors.Validation("invalid registration fields", map[string]string{
			"requirements": "valid email, password >= 6 chars, and non-empty nickname are required",
		})
	}

	_, err := s.playerRepo.FindByEmail(ctx, email)
	if err == nil {
		return "", nil, apperrors.Conflict("email is already registered")
	}
	if !errors.Is(err, usersDomain.ErrPlayerNotFound) {
		return "", nil, apperrors.Internal("failed to verify email availability")
	}

	hash, err := security.HashPassword(password)
	if err != nil {
		return "", nil, apperrors.Internal("failed to process password")
	}

	player, err := usersDomain.NewPlayer(email, hash, nickname)
	if err != nil {
		return "", nil, apperrors.Validation(err.Error(), nil)
	}

	if err := s.playerRepo.Create(ctx, player); err != nil {
		return "", nil, apperrors.Internal("failed to create player account")
	}

	token, err := s.jwtMgr.GeneratePlayerToken(player.ID.String(), string(player.Role))
	if err != nil {
		return "", nil, apperrors.Internal("failed to generate authentication token")
	}

	return token, player, nil
}

// Login autentica a un jugador registrado por correo y contrasena.
func (s *AuthService) Login(ctx context.Context, email, password string) (string, *usersDomain.Player, error) {
	email = strings.ToLower(strings.TrimSpace(email))
	if email == "" || password == "" {
		return "", nil, apperrors.Validation("email and password are required", nil)
	}

	player, err := s.playerRepo.FindByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, usersDomain.ErrPlayerNotFound) {
			return "", nil, apperrors.Unauthorized("invalid email or password")
		}
		return "", nil, apperrors.Internal("failed to fetch player account")
	}

	if err := security.CheckPassword(player.Password, password); err != nil {
		return "", nil, apperrors.Unauthorized("invalid email or password")
	}

	token, err := s.jwtMgr.GeneratePlayerToken(player.ID.String(), string(player.Role))
	if err != nil {
		return "", nil, apperrors.Internal("failed to generate authentication token")
	}

	return token, player, nil
}

// Guest emite un JWT temporal para un jugador invitado sin persistir en PostgreSQL.
func (s *AuthService) Guest(nickname string) (string, *GuestSession, error) {
	nickname = strings.TrimSpace(nickname)
	guestUUID := uuid.New()
	if nickname == "" {
		suffix := guestUUID.ID()%9000 + 1000
		nickname = fmt.Sprintf("Invitado_%d", suffix)
	}

	guestID := guestUUID.String()
	token, err := s.jwtMgr.GenerateGuestToken(guestID, nickname)
	if err != nil {
		return "", nil, apperrors.Internal("failed to generate guest token")
	}

	return token, &GuestSession{
		GuestID:  guestID,
		Nickname: nickname,
		Role:     "guest",
	}, nil
}
