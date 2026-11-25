package infrastructure

import (
	"testing"

	"github.com/google/uuid"
	"github.com/toby-minigames/server/internal/modules/reports/domain"
)

func TestReportMapper_RoundTrip(t *testing.T) {
	gameID := 1
	rep, err := domain.NewReport(uuid.New(), &gameID, "Audio glitch on level 2")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	model := FromDomainReport(rep)
	restored := model.ToDomain()

	if restored.ID != rep.ID || restored.Status != domain.ReportStatusPending || restored.GameID == nil || *restored.GameID != 1 {
		t.Errorf("report round-trip mismatch: got %+v", restored)
	}
}
