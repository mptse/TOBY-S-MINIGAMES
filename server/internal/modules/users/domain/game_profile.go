package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// GameProfile almacena la progresion del jugador por minijuego y su estado dinamico.
type GameProfile struct {
	ID            uuid.UUID
	PlayerID      uuid.UUID
	GameID        int
	HighestScore  int
	MatchesPlayed int
	GameData      map[string]any
	UpdatedAt     time.Time
}

// NewGameProfile inicializa un nuevo subperfil para un jugador y minijuego.
func NewGameProfile(playerID uuid.UUID, gameID int) (*GameProfile, error) {
	if playerID == uuid.Nil {
		return nil, errors.New("player ID cannot be nil")
	}
	if gameID <= 0 {
		return nil, errors.New("game ID must be positive")
	}

	return &GameProfile{
		ID:            uuid.New(),
		PlayerID:      playerID,
		GameID:        gameID,
		HighestScore:  0,
		MatchesPlayed: 0,
		GameData:      map[string]any{},
		UpdatedAt:     time.Now().UTC(),
	}, nil
}
