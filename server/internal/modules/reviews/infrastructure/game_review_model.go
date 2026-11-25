package infrastructure

import (
	"time"

	"github.com/google/uuid"
)

// GameReviewModel es el modelo de persistencia GORM mapeado a la tabla game_reviews.
type GameReviewModel struct {
	ReviewID  uuid.UUID `gorm:"column:review_id;type:uuid;primaryKey"`
	PlayerID  uuid.UUID `gorm:"column:player_id;type:uuid;not null"`
	GameID    int       `gorm:"column:game_id;not null"`
	Rating    int       `gorm:"column:rating;not null"`
	Comment   *string   `gorm:"column:comment;type:text"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamptz;not null"`
}

// TableName retorna el nombre de la tabla en PostgreSQL para GameReviewModel.
func (GameReviewModel) TableName() string {
	return "game_reviews"
}
