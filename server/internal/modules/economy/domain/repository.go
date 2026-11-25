package domain

import (
	"context"

	"github.com/google/uuid"
)

// ItemRepository define las operaciones de persistencia para los objetos de tienda.
type ItemRepository interface {
	Create(ctx context.Context, item *Item) error
	FindByID(ctx context.Context, id uuid.UUID) (*Item, error)
	ListAll(ctx context.Context) ([]*Item, error)
}

// InventoryRepository define las operaciones de persistencia para inventarios de jugadores.
type InventoryRepository interface {
	Add(ctx context.Context, entry *Inventory) error
	ListByPlayer(ctx context.Context, playerID uuid.UUID) ([]*Inventory, error)
	SetEquipped(ctx context.Context, playerID, itemID uuid.UUID, equipped bool) error
}
