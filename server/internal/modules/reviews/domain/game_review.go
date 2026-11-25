package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// GameReview representa la calificacion y comentario opcional de un jugador sobre un minijuego.
type GameReview struct {
	ID        uuid.UUID
	PlayerID  uuid.UUID
	GameID    int
	Rating    int
	Comment   *string
	CreatedAt time.Time
}

// NewGameReview crea y valida una entidad GameReview con calificacion entre 1 y 5.
func NewGameReview(playerID uuid.UUID, gameID, rating int, comment *string) (*GameReview, error) {
	if playerID == uuid.Nil {
		return nil, errors.New("player ID cannot be nil")
	}
	if gameID <= 0 {
		return nil, errors.New("game ID must be positive")
	}
	if rating < 1 || rating > 5 {
		return nil, errors.New("rating must be between 1 and 5")
	}

	return &GameReview{
		ID:        uuid.New(),
		PlayerID:  playerID,
		GameID:    gameID,
		Rating:    rating,
		Comment:   comment,
		CreatedAt: time.Now().UTC(),
	}, nil
}
