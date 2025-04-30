package auth

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	InvalidPasswordLength Error = "e-mar404/auth: password should be at least 8 characters long"
	InvalidPasswordComplexity Error = "e-mar404/auth: password should contain at least one number, symbol and capital case character"
	IncorrectPassword Error = "e-mar404/auth: Incorrect Password"
)
