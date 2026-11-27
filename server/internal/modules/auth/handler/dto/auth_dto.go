package dto

import (
	"github.com/google/uuid"
	usersDomain "github.com/toby-minigames/server/internal/modules/users/domain"
)

// RegisterRequest representa el payload de entrada para registrar un nuevo jugador.
type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Nickname string `json:"nickname" binding:"required"`
}

// LoginRequest representa el payload de entrada para iniciar sesion.
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// GuestRequest representa el payload opcional para crear una sesion de invitado.
type GuestRequest struct {
	Nickname string `json:"nickname"`
}

// PlayerSummaryResponse representa el resumen publico del jugador tras autenticar.
type PlayerSummaryResponse struct {
	PlayerID   uuid.UUID `json:"player_id"`
	Email      string    `json:"email"`
	Nickname   string    `json:"nickname"`
	Role       string    `json:"role"`
	Exp        int       `json:"exp"`
	Tobycoin   int       `json:"tobycoin"`
	AvatarPath string    `json:"avatar_path"`
}

// FromDomainPlayerSummary convierte una entidad Player de dominio a su DTO de respuesta.
func FromDomainPlayerSummary(p *usersDomain.Player) PlayerSummaryResponse {
	return PlayerSummaryResponse{
		PlayerID:   p.ID,
		Email:      p.Email,
		Nickname:   p.Nickname,
		Role:       string(p.Role),
		Exp:        p.Exp,
		Tobycoin:   p.Tobycoin,
		AvatarPath: p.AvatarPath,
	}
}
