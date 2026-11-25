package domain

import "context"

// GameRepository define las operaciones de persistencia para el catalogo de minijuegos.
type GameRepository interface {
	Create(ctx context.Context, game *Game) error
	FindByID(ctx context.Context, id int) (*Game, error)
	FindBySlug(ctx context.Context, slug string) (*Game, error)
	ListActive(ctx context.Context) ([]*Game, error)
}

// CategoryRepository define las operaciones de persistencia para las categorias de minijuegos.
type CategoryRepository interface {
	Create(ctx context.Context, category *Category) error
	ListAll(ctx context.Context) ([]*Category, error)
}
