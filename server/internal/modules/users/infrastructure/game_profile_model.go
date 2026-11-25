package infrastructure

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

// GameProfileModel es el modelo de persistencia GORM mapeado a la tabla game_profiles.
type GameProfileModel struct {
	ProfileID     uuid.UUID      `gorm:"column:profile_id;type:uuid;primaryKey"`
	PlayerID      uuid.UUID      `gorm:"column:player_id;type:uuid;not null"`
	GameID        int            `gorm:"column:game_id;not null"`
	HighestScore  int            `gorm:"column:highest_score;not null;default:0"`
	MatchesPlayed int            `gorm:"column:matches_played;not null;default:0"`
	GameData      datatypes.JSON `gorm:"column:game_data;type:jsonb;not null;default:'{}'"`
	UpdatedAt     time.Time      `gorm:"column:updated_at;type:timestamptz;not null"`
}

// TableName retorna el nombre de la tabla en PostgreSQL para GameProfileModel.
func (GameProfileModel) TableName() string {
	return "game_profiles"
}
