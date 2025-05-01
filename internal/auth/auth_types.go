package auth

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	IncorrectPassword Error = "e-mar404/auth: Incorrect Password"
)
