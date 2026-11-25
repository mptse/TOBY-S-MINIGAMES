package infrastructure

import (
	"github.com/toby-minigames/server/internal/modules/reports/domain"
)

// ToDomain convierte un ReportModel en una entidad pura de dominio Report.
func (m *ReportModel) ToDomain() *domain.Report {
	return &domain.Report{
		ID:          m.ReportID,
		PlayerID:    m.PlayerID,
		GameID:      m.GameID,
		Description: m.Description,
		Status:      domain.ReportStatus(m.Status),
		CreatedAt:   m.CreatedAt,
	}
}

// FromDomainReport convierte una entidad Report de dominio en un ReportModel de GORM.
func FromDomainReport(r *domain.Report) *ReportModel {
	return &ReportModel{
		ReportID:    r.ID,
		PlayerID:    r.PlayerID,
		GameID:      r.GameID,
		Description: r.Description,
		Status:      string(r.Status),
		CreatedAt:   r.CreatedAt,
	}
}
