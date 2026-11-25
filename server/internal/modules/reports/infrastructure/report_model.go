package infrastructure

import (
	"time"

	"github.com/google/uuid"
)

// ReportModel es el modelo de persistencia GORM mapeado a la tabla reports.
type ReportModel struct {
	ReportID    uuid.UUID `gorm:"column:report_id;type:uuid;primaryKey"`
	PlayerID    uuid.UUID `gorm:"column:player_id;type:uuid;not null"`
	GameID      *int      `gorm:"column:game_id"`
	Description string    `gorm:"column:description;type:text;not null"`
	Status      string    `gorm:"column:status;type:varchar(50);not null;default:'pending'"`
	CreatedAt   time.Time `gorm:"column:created_at;type:timestamptz;not null"`
}

// TableName retorna el nombre de la tabla en PostgreSQL para ReportModel.
func (ReportModel) TableName() string {
	return "reports"
}
