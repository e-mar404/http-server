package test

import (
	"e-mar404/http-server/internal/auth"
	"testing"
)

func TestHashPassword(t *testing.T) {
	tt := []struct {
		password string
		attempt  string
		expected error
	}{
		{
			password: "password",
			attempt:  "notpassword",
			expected: auth.IncorrectPassword,
		},
		{
			password: "password",
			attempt:  "password",
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
