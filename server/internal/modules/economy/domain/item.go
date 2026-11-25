package domain

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

// Item representa una recompensa o cosmetico adquirible en la tienda virtual.
type Item struct {
	ID           uuid.UUID
	GameID       *int
	Name         string
	Description  string
	Price        int
	Metadata     map[string]any
	CategorySlug string
}

// Inventory representa un objeto adquirido por un jugador especifico.
type Inventory struct {
	ID         uuid.UUID
	PlayerID   uuid.UUID
	ItemID     uuid.UUID
	IsEquipped bool
}

// NewItem crea y valida una nueva entidad Item de tienda.
func NewItem(name, description, categorySlug string, price int, gameID *int, metadata map[string]any) (*Item, error) {
	name = strings.TrimSpace(name)
	description = strings.TrimSpace(description)
	categorySlug = strings.TrimSpace(categorySlug)

	if name == "" {
		return nil, errors.New("item name cannot be empty")
	}
	if description == "" {
		return nil, errors.New("item description cannot be empty")
	}
	if categorySlug == "" {
		return nil, errors.New("item category slug cannot be empty")
	}
	if price < 0 {
		return nil, errors.New("item price cannot be negative")
	}
	if metadata == nil {
		metadata = map[string]any{}
	}

	return &Item{
		ID:           uuid.New(),
		GameID:       gameID,
		Name:         name,
		Description:  description,
		Price:        price,
		Metadata:     metadata,
		CategorySlug: categorySlug,
	}, nil
}

// NewInventory crea y valida un nuevo registro Inventory para un jugador.
func NewInventory(playerID, itemID uuid.UUID) (*Inventory, error) {
	if playerID == uuid.Nil {
		return nil, errors.New("player ID cannot be nil")
	}
	if itemID == uuid.Nil {
		return nil, errors.New("item ID cannot be nil")
	}

	return &Inventory{
		ID:         uuid.New(),
		PlayerID:   playerID,
		ItemID:     itemID,
		IsEquipped: false,
	}, nil
}
