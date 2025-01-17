package exception

import "errors"

var (
	UserEmailAlreadyExist       = errors.New("email already exist")
	UserEmailAndPasswordIsWrong = errors.New("email or password is wrong")
	UserNotFound                = errors.New("user not found")
)
