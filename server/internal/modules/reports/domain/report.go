package domain

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

// ReportStatus representa el estado del ciclo de vida de un reporte enviado.
type ReportStatus string

const (
	ReportStatusPending   ReportStatus = "pending"
	ReportStatusInReview  ReportStatus = "in_review"
	ReportStatusResolved  ReportStatus = "resolved"
	ReportStatusDismissed ReportStatus = "dismissed"
)

// Report representa un reporte de fallo tecnico o comportamiento enviado por un jugador.
type Report struct {
	ID          uuid.UUID
	PlayerID    uuid.UUID
	GameID      *int
	Description string
	Status      ReportStatus
	CreatedAt   time.Time
}

// NewReport crea y valida una nueva entidad Report con estado inicial pending.
func NewReport(playerID uuid.UUID, gameID *int, description string) (*Report, error) {
	if playerID == uuid.Nil {
		return nil, errors.New("player ID cannot be nil")
	}
	description = strings.TrimSpace(description)
	if description == "" {
		return nil, errors.New("report description cannot be empty")
	}

	return &Report{
		ID:          uuid.New(),
		PlayerID:    playerID,
		GameID:      gameID,
		Description: description,
		Status:      ReportStatusPending,
		CreatedAt:   time.Now().UTC(),
	}, nil
}
