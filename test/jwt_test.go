package test

import (
	"crypto/rand"
	"e-mar404/http-server/internal/auth"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func TestMakeJWTAndValidateJWT(t *testing.T) {
	userID := uuid.New()
	secret := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, secret)
	if err != nil {
		t.Fatalf("unable to make a secret: %v", err)
	}

	tt := []struct{
		tokenSecret string
		expiresIn time.Duration
		waitFor time.Duration
		expectedErr error
	}{
		{
			tokenSecret: string(secret),
			expiresIn: 1 * time.Second,
			waitFor: 2 * time.Second,
			expectedErr: jwt.ErrTokenExpired,
		},
		{
			tokenSecret: "not the correct token secret",
			expiresIn: 2 * time.Second,
			waitFor: 1 * time.Second,
			expectedErr: jwt.ErrTokenSignatureInvalid, 
		},
		{
			tokenSecret: string(secret),
			expiresIn: 2 * time.Second,
			waitFor: 1 * time.Second,
			expectedErr: nil, 
		},
	}

	for _, tc := range tt {
		token, err := auth.MakeJWT(userID, string(secret), tc.expiresIn)
		time.Sleep(tc.waitFor)
		if err != nil {
			t.Fatalf("unexpected error when making jwt: %v", err)
		}

		id, err := auth.ValidateJWT(token, tc.tokenSecret)
		if err != nil {
			if !strings.Contains(err.Error(), tc.expectedErr.Error()) {
				t.Fatalf("Validation errors do not match. Expected: %v, Got: %v", tc.expectedErr, err)
			}
		}

		if userID != id && tc.expectedErr == nil {
			t.Fatalf("user ids do not match. Expected: %v, Got: %v", userID, id)
		}
	}
}
