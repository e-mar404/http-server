package test

import (
	"e-mar404/http-server/internal/auth"
	"testing"
)

func TestHashPassword(t *testing.T) {
	tt := []struct {
		password string
		attempt string
		expected error
	}{
		{
			password : "hello",
			attempt: "",
			expected: auth.InvalidPasswordLength,
		},
		{
			password: "thisislongenough",
			attempt: "",
			expected: auth.InvalidPasswordComplexity,
		},
		{
			password: "Th!sW0rks",
			attempt: "Th!sDoesNotW0rks",
			expected: auth.IncorrectPassword,
		},
		{
			password: "Th!sW0rks",
			attempt: "Th!sW0rks",
			expected: nil,
		},
	}

	for _, tc := range tt {
		hash, hasErr := auth.HashPassword(tc.password)
		if tc.attempt == "" && hasErr != tc.expected {
			t.Fatalf("hash password: actual error did not match expected, Expected: %v Got: %v", tc.expected, hasErr)
		}

		checkErr := auth.CheckPasswordHash(hash, tc.attempt)
		if tc.attempt != "" && checkErr != tc.expected {
			t.Fatalf("check password: actual error did not match expected, Expected: %v Got: %v", tc.expected, checkErr)
		}
	}
}
