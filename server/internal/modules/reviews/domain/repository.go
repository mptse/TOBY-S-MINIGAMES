package domain

import "context"

// GameReviewRepository define las operaciones de persistencia para las resenas de juegos.
type GameReviewRepository interface {
	Create(ctx context.Context, review *GameReview) error
	ListByGame(ctx context.Context, gameID, limit, offset int) ([]*GameReview, error)
}
