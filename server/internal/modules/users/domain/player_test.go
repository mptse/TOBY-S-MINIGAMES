package domain

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewPlayer_Success(t *testing.T) {
	p, err := NewPlayer("toby@example.com", "hashed_pw", "Toby")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if p.ID == uuid.Nil {
		t.Error("expected generated non-nil UUID")
	}
	if p.Role != RolePlayer {
		t.Errorf("expected role %s, got %s", RolePlayer, p.Role)
	}
}

func TestNewPlayer_ValidationErrors(t *testing.T) {
	if _, err := NewPlayer("", "hash", "Toby"); err == nil {
		t.Error("expected error for empty email")
	}
	if _, err := NewPlayer("a@b.com", "", "Toby"); err == nil {
		t.Error("expected error for empty password hash")
	}
	if _, err := NewPlayer("a@b.com", "hash", "   "); err == nil {
		t.Error("expected error for empty nickname")
	}
}

func TestPlayer_CoinsAndRewards(t *testing.T) {
	p, _ := NewPlayer("toby@example.com", "hash", "Toby")
	if err := p.AddRewards(100, 50); err != nil {
		t.Fatalf("unexpected error adding rewards: %v", err)
	}
	if p.Exp != 100 || p.Tobycoin != 50 {
		t.Errorf("expected exp=100, coins=50, got exp=%d, coins=%d", p.Exp, p.Tobycoin)
	}
	if err := p.DeductCoins(60); err == nil {
		t.Error("expected error for insufficient balance")
	}
	if err := p.DeductCoins(30); err != nil {
		t.Fatalf("unexpected error deducting coins: %v", err)
	}
	if p.Tobycoin != 20 {
		t.Errorf("expected remaining coins=20, got %d", p.Tobycoin)
	}
}

func TestNewGameProfile_Validation(t *testing.T) {
	if _, err := NewGameProfile(uuid.Nil, 1); err == nil {
		t.Error("expected error for nil player ID")
	}
	if _, err := NewGameProfile(uuid.New(), 0); err == nil {
		t.Error("expected error for non-positive game ID")
	}
	gp, err := NewGameProfile(uuid.New(), 1)
	if err != nil {
		t.Fatalf("unexpected error creating game profile: %v", err)
	}
	if gp.GameData == nil {
		t.Error("expected initialized non-nil GameData map")
	}
}
