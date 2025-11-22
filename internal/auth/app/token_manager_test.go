package auth

import (
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestTokenManager_GenerateToken(t *testing.T) {
	secret := os.Getenv("SECRET_JWT")
	tm := NewTokenManager(secret, time.Minute)
	jwt, err := tm.GenerateToken(uuid.NewString())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if jwt == "" {
		t.Fatal("expected token string, got empty")
	}

}

func TestTokenManager_ValidateToken(t *testing.T) {
	secret := os.Getenv("SECRET_JWT")
	userId := uuid.NewString()
	tm := NewTokenManager(secret, time.Minute)
	jwt, err := tm.GenerateToken(userId)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if jwt == "" {
		t.Fatal("expected token string, got empty")
	}
	sub, err := tm.ValidateToken(jwt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if sub == "" {
		t.Fatal("expected sub string, got empty")
	}

	if sub != userId {
		t.Fatalf("expected sub=%s, got %s", userId, sub)
	}
}
