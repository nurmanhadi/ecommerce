package exception

import "errors"

type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

var (
	UserEmailAlreadyExist       = errors.New("email already exist")
	UserEmailAndPasswordIsWrong = errors.New("email or password is wrong")
)
