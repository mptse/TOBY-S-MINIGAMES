package security

import (
	"testing"
)

func TestPasswordHashing(t *testing.T) {
	hash, err := HashPassword("example")
	if err != nil {
		t.Fatalf("unexpected hashing error: %v", err)
	}
	if err := CheckPassword(hash, "example"); err != nil {
		t.Errorf("expected valid password match, got %v", err)
	}
	if err := CheckPassword(hash, "wrong"); err == nil {
		t.Error("expected error for mismatched password")
	}
}

func TestJWTManager_PlayerToken(t *testing.T) {
	mgr, err := NewJWTManager("test", 24)
	if err != nil {
		t.Fatalf("unexpected error creating JWTManager: %v", err)
	}

	playerToken, err := mgr.GeneratePlayerToken("player-uuid-1", "player")
	if err != nil {
		t.Fatalf("unexpected error generating player token: %v", err)
	}

	claims, err := mgr.VerifyToken(playerToken)
	if err != nil {
		t.Fatalf("failed verifying player token: %v", err)
	}
	if claims.Subject != "player-uuid-1" || claims.Role != "player" || claims.IsGuest {
		t.Errorf("unexpected player claims: %+v", claims)
	}
}

func TestJWTManager_GuestToken(t *testing.T) {
	mgr, err := NewJWTManager("test", 24)
	if err != nil {
		t.Fatalf("unexpected error creating JWTManager: %v", err)
	}

	guestToken, err := mgr.GenerateGuestToken("guest-uuid-1", "Invitado_1234")
	if err != nil {
		t.Fatalf("unexpected error generating guest token: %v", err)
	}

	guestClaims, err := mgr.VerifyToken(guestToken)
	if err != nil {
		t.Fatalf("failed verifying guest token: %v", err)
	}
	if !guestClaims.IsGuest || guestClaims.Role != "guest" || guestClaims.Nickname != "Invitado_1234" {
		t.Errorf("unexpected guest claims: %+v", guestClaims)
	}
}
