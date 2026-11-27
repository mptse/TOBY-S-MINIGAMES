package dto

import (
	"time"

	"github.com/google/uuid"
	usersDomain "github.com/toby-minigames/server/internal/modules/users/domain"
)

// UpdateMeRequest representa el payload de actualizacion parcial del perfil propio.
type UpdateMeRequest struct {
	Nickname   *string `json:"nickname"`
	Bio        *string `json:"bio"`
	AvatarPath *string `json:"avatar_path"`
}

// PlayerResponse representa la vista completa del perfil del jugador autenticado.
type PlayerResponse struct {
	PlayerID   uuid.UUID `json:"player_id"`
	Email      string    `json:"email"`
	Nickname   string    `json:"nickname"`
	Bio        *string   `json:"bio"`
	Role       string    `json:"role"`
	Exp        int       `json:"exp"`
	Tobycoin   int       `json:"tobycoin"`
	AvatarPath string    `json:"avatar_path"`
}

// GameProfileResponse representa el subperfil de un jugador para un minijuego.
type GameProfileResponse struct {
	ProfileID     uuid.UUID      `json:"profile_id"`
	GameID        int            `json:"game_id"`
	HighestScore  int            `json:"highest_score"`
	MatchesPlayed int            `json:"matches_played"`
	GameData      map[string]any `json:"game_data"`
	UpdatedAt     time.Time      `json:"updated_at"`
}

// FromDomainPlayer convierte una entidad Player de dominio en su DTO de respuesta.
func FromDomainPlayer(p *usersDomain.Player) PlayerResponse {
	return PlayerResponse{
		PlayerID:   p.ID,
		Email:      p.Email,
		Nickname:   p.Nickname,
		Bio:        p.Bio,
		Role:       string(p.Role),
		Exp:        p.Exp,
		Tobycoin:   p.Tobycoin,
		AvatarPath: p.AvatarPath,
	}
}

// FromDomainGameProfile convierte una entidad GameProfile de dominio en su DTO de respuesta.
func FromDomainGameProfile(gp *usersDomain.GameProfile) GameProfileResponse {
	return GameProfileResponse{
		ProfileID:     gp.ID,
		GameID:        gp.GameID,
		HighestScore:  gp.HighestScore,
		MatchesPlayed: gp.MatchesPlayed,
		GameData:      gp.GameData,
		UpdatedAt:     gp.UpdatedAt,
	}
}
