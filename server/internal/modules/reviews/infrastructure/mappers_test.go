package infrastructure

import (
	"testing"

	"github.com/google/uuid"
	"github.com/toby-minigames/server/internal/modules/reviews/domain"
)

func TestGameReviewMapper_RoundTrip(t *testing.T) {
	comment := "Fun and responsive"
	review, err := domain.NewGameReview(uuid.New(), 5, 4, &comment)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	model := FromDomainGameReview(review)
	restored := model.ToDomain()

	if restored.ID != review.ID || restored.Rating != 4 || *restored.Comment != comment {
		t.Errorf("game review round-trip mismatch: got %+v", restored)
	}
}
