package exception

import "errors"

var (
	JwtInvalidSignature = errors.New("token signature is invalid")
	JwtExpired          = errors.New("token has expired")
)
