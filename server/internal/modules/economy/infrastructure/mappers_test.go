package infrastructure

import (
	"testing"

	"github.com/google/uuid"
	"github.com/toby-minigames/server/internal/modules/economy/domain"
)

func TestItemMapper_RoundTripWithJSONB(t *testing.T) {
	gameID := 4
	item, err := domain.NewItem("Neon Skin", "Glowing avatar skin", "cosmetic", 250, &gameID, map[string]any{
		"rarity": "epic",
		"color":  "#00FFAA",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	model, err := FromDomainItem(item)
	if err != nil {
		t.Fatalf("failed mapping item to model: %v", err)
	}

	restored, err := model.ToDomain()
	if err != nil {
		t.Fatalf("failed mapping model to domain: %v", err)
	}

	if restored.ID != item.ID || restored.Name != "Neon Skin" || restored.GameID == nil || *restored.GameID != 4 {
		t.Errorf("item scalar mismatch: got %+v", restored)
	}
	if restored.Metadata["rarity"] != "epic" {
		t.Errorf("expected rarity='epic', got %v", restored.Metadata["rarity"])
	}
}

func TestInventoryMapper_RoundTrip(t *testing.T) {
	inv, err := domain.NewInventory(uuid.New(), uuid.New())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	inv.IsEquipped = true

	model := FromDomainInventory(inv)
	restored := model.ToDomain()

	if restored.ID != inv.ID || !restored.IsEquipped {
		t.Errorf("inventory round-trip mismatch: got %+v", restored)
	}
}
