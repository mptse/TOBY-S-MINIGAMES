package domain

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

// Errores de dominio para repositorios del modulo users.
var (
	ErrPlayerNotFound      = errors.New("player not found")
	ErrGameProfileNotFound = errors.New("game profile not found")
)

// PlayerRepository define las operaciones de persistencia para entidades Player.
type PlayerRepository interface {
	Create(ctx context.Context, player *Player) error
	FindByID(ctx context.Context, id uuid.UUID) (*Player, error)
	FindByEmail(ctx context.Context, email string) (*Player, error)
	Update(ctx context.Context, player *Player) error
}

// GameProfileRepository define las operaciones de persistencia para subperfiles por juego.
type GameProfileRepository interface {
	Upsert(ctx context.Context, profile *GameProfile) error
	FindByPlayerAndGame(ctx context.Context, playerID uuid.UUID, gameID int) (*GameProfile, error)
	ListByPlayer(ctx context.Context, playerID uuid.UUID) ([]*GameProfile, error)
}
