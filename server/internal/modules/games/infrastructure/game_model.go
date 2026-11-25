package infrastructure

import (
	"time"
)

// GameModel es el modelo de persistencia GORM mapeado a la tabla games.
type GameModel struct {
	GameID        int             `gorm:"column:game_id;primaryKey;autoIncrement"`
	Name          string          `gorm:"column:name;type:varchar(150);not null"`
	NameSlug      string          `gorm:"column:name_slug;type:varchar(150);uniqueIndex;not null"`
	Description   string          `gorm:"column:description;type:varchar(1000);not null"`
	ThumbnailSlug string          `gorm:"column:thumbnail_slug;type:varchar(255);not null"`
	IsActive      bool            `gorm:"column:is_active;not null;default:true"`
	CreateAt      time.Time       `gorm:"column:create_at;type:timestamptz;not null"`
	Categories    []CategoryModel `gorm:"many2many:game_categories;joinForeignKey:game_id;joinReferences:category_id"`
}

// TableName retorna el nombre de la tabla en PostgreSQL para GameModel.
func (GameModel) TableName() string {
	return "games"
}

// CategoryModel es el modelo de persistencia GORM mapeado a la tabla categories.
type CategoryModel struct {
	CategoryID   int     `gorm:"column:category_id;primaryKey;autoIncrement"`
	NameCategory string  `gorm:"column:name_category;type:varchar(100);not null"`
	CategoryDesc *string `gorm:"column:category_desc;type:varchar(500)"`
	CategorySlug string  `gorm:"column:category_slug;type:varchar(100);uniqueIndex;not null"`
}

// TableName retorna el nombre de la tabla en PostgreSQL para CategoryModel.
func (CategoryModel) TableName() string {
	return "categories"
}

// GameCategoryModel es el modelo GORM de la tabla intermedia game_categories.
type GameCategoryModel struct {
	GameID     int `gorm:"column:game_id;primaryKey"`
	CategoryID int `gorm:"column:category_id;primaryKey"`
}

// TableName retorna el nombre de la tabla en PostgreSQL para GameCategoryModel.
func (GameCategoryModel) TableName() string {
	return "game_categories"
}
