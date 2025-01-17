package exception

import "errors"

var (
	AuthTokenIsNull   = errors.New("token authorization null")
	AuthBearerReqired = errors.New("need token jwt bearer example 'Bearer oqieuoqeiq'")
)
