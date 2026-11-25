package domain

import (
	"context"

	"github.com/google/uuid"
)

// MatchHistoryRepository define las operaciones de persistencia para el historial de partidas.
type MatchHistoryRepository interface {
	Create(ctx context.Context, match *MatchHistory) error
	ListByPlayer(ctx context.Context, playerID uuid.UUID, limit, offset int) ([]*MatchHistory, error)
}
