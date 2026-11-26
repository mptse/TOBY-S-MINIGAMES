package database

import (
	"context"
	"testing"
)

func TestNewConnection_EmptyURL(t *testing.T) {
	_, err := NewConnection("", "development", DefaultPoolConfig())
	if err == nil {
		t.Fatal("expected an error for empty databaseURL, got nil")
	}
}

func TestPing_NilDB(t *testing.T) {
	err := Ping(context.Background(), nil)
	if err == nil {
		t.Fatal("expected an error for nil database handle, got nil")
	}
}
