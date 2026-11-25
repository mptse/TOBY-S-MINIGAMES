package infrastructure

import (
	"testing"

	"github.com/google/uuid"
	"github.com/toby-minigames/server/internal/modules/scores/domain"
)

func TestMatchHistoryMapper_RoundTrip(t *testing.T) {
	mh, err := domain.NewMatchHistory(uuid.New(), 2, "timed", 1500, 75, 15)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	model := FromDomainMatchHistory(mh)
	restored := model.ToDomain()

	if restored.ID != mh.ID || restored.ScoreObtained != 1500 || restored.GameMode != "timed" {
		t.Errorf("match history round-trip mismatch: got %+v", restored)
	}
}
