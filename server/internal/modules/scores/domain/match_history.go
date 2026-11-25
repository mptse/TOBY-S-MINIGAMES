package domain

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

// MatchHistory registra el resultado y las recompensas obtenidas en una partida finalizada.
type MatchHistory struct {
	ID              uuid.UUID
	PlayerID        uuid.UUID
	GameID          int
	GameMode        string
	ScoreObtained   int
	XPEarned        int
	TobycoinsEarned int
	PlayedAt        time.Time
}

// NewMatchHistory crea y valida un nuevo registro MatchHistory.
func NewMatchHistory(playerID uuid.UUID, gameID int, gameMode string, score, xpEarned, coinsEarned int) (*MatchHistory, error) {
	if playerID == uuid.Nil {
		return nil, errors.New("player ID cannot be nil")
	}
	if gameID <= 0 {
		return nil, errors.New("game ID must be positive")
	}
	gameMode = strings.TrimSpace(gameMode)
	if gameMode == "" {
		gameMode = "standard"
	}
	if xpEarned < 0 || coinsEarned < 0 {
		return nil, errors.New("earned XP and tobycoins cannot be negative")
	}

	return &MatchHistory{
		ID:              uuid.New(),
		PlayerID:        playerID,
		GameID:          gameID,
		GameMode:        gameMode,
		ScoreObtained:   score,
		XPEarned:        xpEarned,
		TobycoinsEarned: coinsEarned,
		PlayedAt:        time.Now().UTC(),
	}, nil
}
