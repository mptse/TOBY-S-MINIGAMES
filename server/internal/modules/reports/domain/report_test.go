package domain

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewReport_Validation(t *testing.T) {
	if _, err := NewReport(uuid.Nil, nil, "Bug found"); err == nil {
		t.Error("expected error for nil player ID")
	}
	if _, err := NewReport(uuid.New(), nil, "   "); err == nil {
		t.Error("expected error for empty description")
	}

	gameID := 2
	rep, err := NewReport(uuid.New(), &gameID, "Score not saving")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if rep.Status != ReportStatusPending {
		t.Errorf("expected initial status %s, got %s", ReportStatusPending, rep.Status)
	}
}
