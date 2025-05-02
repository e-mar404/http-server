package auth

import (
	"fmt"
	"net/http"
	"strings"
)

func GetBearerToken(headers  http.Header) (string, error) {
	bearerToken := headers.Get("Authorization")
	if bearerToken == "" {
		return "", fmt.Errorf("Missing Bearer token value from Authorization header")
	}

	token, found := strings.CutPrefix(bearerToken, "Bearer ")
	if !found {
		return "", fmt.Errorf("Missing `Bearer ` prefix")
	}

	return token, nil
}
