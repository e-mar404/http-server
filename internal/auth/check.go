package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func CheckPasswordHash(hash, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return IncorrectPassword
	}
	return nil
}
