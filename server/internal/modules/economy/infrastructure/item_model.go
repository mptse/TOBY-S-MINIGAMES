package infrastructure

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

// ItemModel es el modelo de persistencia GORM mapeado a la tabla items.
type ItemModel struct {
	ItemID        uuid.UUID      `gorm:"column:item_id;type:uuid;primaryKey"`
	GameID        *int           `gorm:"column:game_id"`
	ItemName      string         `gorm:"column:item_name;type:varchar(150);uniqueIndex;not null"`
	DescItem      string         `gorm:"column:desc_item;type:varchar(500);not null"`
	Price         int            `gorm:"column:price;not null"`
	Metadata      datatypes.JSON `gorm:"column:metadata;type:jsonb;not null;default:'{}'"`
	CategoriaSlug string         `gorm:"column:categoria_slug;type:varchar(100);not null"`
}

// TableName retorna el nombre de la tabla en PostgreSQL para ItemModel.
func (ItemModel) TableName() string {
	return "items"
}

// InventoryModel es el modelo de persistencia GORM mapeado a la tabla inventories.
type InventoryModel struct {
	InventoryID uuid.UUID `gorm:"column:inventory_id;type:uuid;primaryKey"`
	PlayerID    uuid.UUID `gorm:"column:player_id;type:uuid;not null"`
	ItemID      uuid.UUID `gorm:"column:item_id;type:uuid;not null"`
	IsEquiped   bool      `gorm:"column:is_equiped;not null;default:false"`
}

// TableName retorna el nombre de la tabla en PostgreSQL para InventoryModel.
func (InventoryModel) TableName() string {
	return "inventories"
}
