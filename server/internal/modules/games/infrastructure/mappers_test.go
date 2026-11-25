package infrastructure

import (
	"testing"

	"github.com/toby-minigames/server/internal/modules/games/domain"
)

func TestGameMapper_RoundTrip(t *testing.T) {
	desc := "Fast-paced questions"
	g, err := domain.NewGame("Trivia Master", "trivia-master", "Test your knowledge", "trivia-thumb.png")
	if err != nil {
		t.Fatalf("unexpected error creating game: %v", err)
	}
	g.ID = 7
	g.Categories = []domain.Category{
		{ID: 1, Name: "Trivia", Description: &desc, Slug: "trivia"},
	}

	model := FromDomainGame(g)
	restored := model.ToDomain()

	if restored.ID != 7 || restored.Slug != "trivia-master" {
		t.Errorf("game scalar mismatch: got %+v", restored)
	}
	if len(restored.Categories) != 1 || restored.Categories[0].Slug != "trivia" {
		t.Errorf("expected 1 category with slug 'trivia', got %+v", restored.Categories)
	}
}
