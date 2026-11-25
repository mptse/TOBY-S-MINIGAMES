package domain

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewItem_Validation(t *testing.T) {
	if _, err := NewItem("", "desc", "avatar", 100, nil, nil); err == nil {
		t.Error("expected error for empty item name")
	}
	if _, err := NewItem("Gold Crown", "desc", "avatar", -10, nil, nil); err == nil {
		t.Error("expected error for negative price")
	}

	item, err := NewItem("Gold Crown", "Shiny crown", "avatar", 500, nil, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if item.Metadata == nil {
		t.Error("expected initialized non-nil metadata map")
	}
}

func TestNewInventory_Validation(t *testing.T) {
	if _, err := NewInventory(uuid.Nil, uuid.New()); err == nil {
		t.Error("expected error for nil player ID")
	}
	if _, err := NewInventory(uuid.New(), uuid.Nil); err == nil {
		t.Error("expected error for nil item ID")
	}
}
