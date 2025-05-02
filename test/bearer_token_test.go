package test

import (
	"e-mar404/http-server/internal/auth"
	"errors"
	"fmt"
	"net/http"
	"testing"
)

func TestGetBearerToken(t *testing.T) {
	tt := []struct {
		header http.Header
		token string
		expectedToken string
		expectedErr error
	}{
		{
			header: http.Header{},
			token: "nobearer",
			expectedToken: "",
			expectedErr: fmt.Errorf("Missing `Bearer ` prefix"),
		},
		{
			header: http.Header{},
			token: "",
			expectedToken: "",
			expectedErr: fmt.Errorf("Missing Bearer token value from Authorization header"),
		},
		{
			header: http.Header{},
			token: "Bearer sometoken",
			expectedToken: "sometoken",
			expectedErr: nil, 
		},
	}

	for _, tc := range tt {
		tc.header.Add("Authorization", tc.token)
		token, err := auth.GetBearerToken(tc.header)
		if tc.expectedErr != nil {
			if errors.Is(err, tc.expectedErr) {
				t.Fatalf("Incorrect error. Expected: %v, Got: %v", tc.expectedErr, err)
			}
		}
		if token != tc.expectedToken && tc.expectedToken != "" {
			t.Fatalf("Incorrect token. Expected: %v, Got: %v", tc.token, token)
		}
	}
}
