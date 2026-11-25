package domain

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewGameReview_Validation(t *testing.T) {
	if _, err := NewGameReview(uuid.New(), 1, 0, nil); err == nil {
		t.Error("expected error for rating < 1")
	}
	if _, err := NewGameReview(uuid.New(), 1, 6, nil); err == nil {
		t.Error("expected error for rating > 5")
	}

	comment := "Great minigame!"
	r, err := NewGameReview(uuid.New(), 1, 5, &comment)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if r.Rating != 5 || r.Comment == nil || *r.Comment != comment {
		t.Errorf("unexpected review values: %+v", r)
	}
}
