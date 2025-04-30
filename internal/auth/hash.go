package auth

import (
	"errors"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	err := checkConstraints(password)
	if err != nil {
		return "", err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	
	return string(hash), nil
}

func checkConstraints(password string) error {
	hasNumber, hasSymbol, hasUppercase := false, false, false
	if len(password) < 8 {
		return errors.New("e-mar404/auth: password should be at least 8 characters long")
	}
	for _, letter := range password {
		if unicode.IsNumber(letter) {
			hasNumber = true
		}
		if unicode.IsSymbol(letter) {
			hasSymbol = true
		}
		if unicode.IsUpper(letter) {
			hasUppercase = true
		}
	}

	if !hasNumber || !hasSymbol || !hasUppercase {
		return errors.New("e-mar404/auth: passwrod should contain at least one number, symbol and capital case character")
	}
	
	return nil
}
