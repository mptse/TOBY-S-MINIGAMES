package infrastructure

import (
	"time"

	"github.com/google/uuid"
)

// MatchHistoryModel es el modelo de persistencia GORM mapeado a la tabla match_history.
type MatchHistoryModel struct {
	MatchID         uuid.UUID `gorm:"column:match_id;type:uuid;primaryKey"`
	PlayerID        uuid.UUID `gorm:"column:player_id;type:uuid;not null"`
	GameID          int       `gorm:"column:game_id;not null"`
	GameMode        string    `gorm:"column:game_mode;type:varchar(100);not null;default:'standard'"`
	ScoreObtained   int       `gorm:"column:score_obtained;not null;default:0"`
	XPEarned        int       `gorm:"column:xp_earned;not null;default:0"`
	TobycoinsEarned int       `gorm:"column:tobycoins_earned;not null;default:0"`
	PlayedAt        time.Time `gorm:"column:played_at;type:timestamptz;not null"`
}

// TableName retorna el nombre de la tabla en PostgreSQL para MatchHistoryModel.
func (MatchHistoryModel) TableName() string {
	return "match_history"
}
