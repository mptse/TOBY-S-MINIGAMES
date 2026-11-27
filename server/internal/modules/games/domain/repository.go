package domain

import (
	"context"
	"errors"
)

// Errores centinela de dominio para el modulo games.
var (
	ErrGameNotFound     = errors.New("game not found")
	ErrCategoryNotFound = errors.New("category not found")
)

// GameRepository define las operaciones de persistencia para el catalogo de minijuegos.
type GameRepository interface {
	Create(ctx context.Context, game *Game) error
	FindByID(ctx context.Context, id int) (*Game, error)
	FindBySlug(ctx context.Context, slug string) (*Game, error)
	ListActive(ctx context.Context) ([]*Game, error)
	ListActiveByCategory(ctx context.Context, categorySlug string) ([]*Game, error)
}

// CategoryRepository define las operaciones de persistencia para las categorias de minijuegos.
type CategoryRepository interface {
	Create(ctx context.Context, category *Category) error
	FindBySlug(ctx context.Context, slug string) (*Category, error)
	ListAll(ctx context.Context) ([]*Category, error)
}
