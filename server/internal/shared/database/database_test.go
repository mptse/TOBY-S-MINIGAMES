package database

import (
	"testing"
)

func TestNewConnection_EmptyURL(t *testing.T) {
	_, err := NewConnection("", DefaultPoolConfig())
	if err == nil {
		t.Fatal("expected an error for empty databaseURL, got nil")
	}
}
