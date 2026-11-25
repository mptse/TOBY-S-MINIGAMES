package infrastructure

import (
	"github.com/toby-minigames/server/internal/modules/economy/domain"
	"github.com/toby-minigames/server/internal/shared/database"
)

// ToDomain convierte un ItemModel en una entidad pura de dominio Item.
func (m *ItemModel) ToDomain() (*domain.Item, error) {
	meta, err := database.JSONBToMap(m.Metadata)
	if err != nil {
		return nil, err
	}

	return &domain.Item{
		ID:           m.ItemID,
		GameID:       m.GameID,
		Name:         m.ItemName,
		Description:  m.DescItem,
		Price:        m.Price,
		Metadata:     meta,
		CategorySlug: m.CategoriaSlug,
	}, nil
}

// FromDomainItem convierte una entidad Item de dominio en un ItemModel de GORM.
func FromDomainItem(item *domain.Item) (*ItemModel, error) {
	jsonMeta, err := database.MapToJSONB(item.Metadata)
	if err != nil {
		return nil, err
	}

	return &ItemModel{
		ItemID:        item.ID,
		GameID:        item.GameID,
		ItemName:      item.Name,
		DescItem:      item.Description,
		Price:         item.Price,
		Metadata:      jsonMeta,
		CategoriaSlug: item.CategorySlug,
	}, nil
}

// ToDomain convierte un InventoryModel en una entidad pura de dominio Inventory.
func (m *InventoryModel) ToDomain() *domain.Inventory {
	return &domain.Inventory{
		ID:         m.InventoryID,
		PlayerID:   m.PlayerID,
		ItemID:     m.ItemID,
		IsEquipped: m.IsEquiped,
	}
}

// FromDomainInventory convierte una entidad Inventory de dominio en un InventoryModel de GORM.
func FromDomainInventory(inv *domain.Inventory) *InventoryModel {
	return &InventoryModel{
		InventoryID: inv.ID,
		PlayerID:    inv.PlayerID,
		ItemID:      inv.ItemID,
		IsEquiped:   inv.IsEquipped,
	}
}
