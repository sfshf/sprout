package schema

import "errors"

var (
	ErrInvalidCaptcha           = errors.New(InvalidCaptcha.String())
	ErrInvalidAccountOrPassword = errors.New(InvalidAccountOrPassword.String())
	ErrUnauthorized             = errors.New(Unauthorized.String())
)
