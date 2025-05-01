package auth

import (
	// "unicode"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	// err := checkConstraints(password)
	// if err != nil {
	// 	return "", err
	// }

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// func checkConstraints(password string) error {
// 	hasNumber, hasSpecial, hasUppercase := false, false, false
// 	if len(password) < 8 {
// 		return InvalidPasswordLength
// 	}
// 	for _, letter := range []rune(password) {
// 		if unicode.IsNumber(letter) {
// 			hasNumber = true
// 		}
// 		if unicode.IsSymbol(letter) || unicode.IsPunct(letter) {
// 			hasSpecial = true
// 		}
// 		if unicode.IsUpper(letter) {
// 			hasUppercase = true
// 		}
// 	}
//
// 	if !hasNumber || !hasSpecial || !hasUppercase {
// 		return InvalidPasswordComplexity
// 	}
//
// 	return nil
// }
