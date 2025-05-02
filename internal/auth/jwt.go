package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func MakeJWT(userID uuid.UUID, tokenSecret string, expiresIn time.Duration) (string, error) {
	currentTime := time.Now()
	claims := jwt.RegisteredClaims{
		Issuer:    "chirpy",
		Subject:   userID.String(),
		ExpiresAt: jwt.NewNumericDate(currentTime.Add(expiresIn)),
		IssuedAt:  jwt.NewNumericDate(currentTime),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&jwt.RegisteredClaims{},
		func(t *jwt.Token) (any, error) {
			return []byte(tokenSecret), nil
		})
	if err != nil {
		return uuid.UUID{}, err
	}

	subject, err := token.Claims.GetSubject()
	if err != nil {
		return uuid.UUID{}, err
	}

	userID, err := uuid.Parse(subject)
	if err != nil {
		return uuid.UUID{}, err
	}

	return userID, nil
}
