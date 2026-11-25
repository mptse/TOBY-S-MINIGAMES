package infrastructure

import (
	"testing"

	"github.com/google/uuid"
	"github.com/toby-minigames/server/internal/modules/users/domain"
)

func TestPlayerMapper_RoundTrip(t *testing.T) {
	bio := "Arcade enthusiast"
	original := &domain.Player{
		ID:         uuid.New(),
		Email:      "player@example.com",
		Password:   "example",
		Nickname:   "TobyGamer",
		Bio:        &bio,
		Role:       domain.RolePlayer,
		Exp:        250,
		Tobycoin:   120,
		AvatarPath: "avatar1.png",
	}

	model := FromDomainPlayer(original)
	restored := model.ToDomain()

	if restored.ID != original.ID || restored.Email != original.Email || restored.Nickname != original.Nickname {
		t.Errorf("player round-trip mismatch: got %+v, want %+v", restored, original)
	}
	if restored.Bio == nil || *restored.Bio != bio {
		t.Errorf("expected bio %q, got %v", bio, restored.Bio)
	}
}

func TestGameProfileMapper_RoundTripWithJSONB(t *testing.T) {
	gp, err := domain.NewGameProfile(uuid.New(), 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	gp.HighestScore = 4200
	gp.MatchesPlayed = 12
	gp.GameData = map[string]any{
		"unlocked_worlds": []any{"forest", "desert"},
		"best_combo":      float64(18),
	}

	model, err := FromDomainGameProfile(gp)
	if err != nil {
		t.Fatalf("failed mapping to model: %v", err)
	}

	restored, err := model.ToDomain()
	if err != nil {
		t.Fatalf("failed mapping back to domain: %v", err)
	}

	if restored.ID != gp.ID || restored.HighestScore != 4200 || restored.MatchesPlayed != 12 {
		t.Errorf("game profile scalar fields mismatch: got %+v", restored)
	}
	if restored.GameData["best_combo"] != float64(18) {
		t.Errorf("expected best_combo=18 in GameData, got %v", restored.GameData["best_combo"])
	}
}
