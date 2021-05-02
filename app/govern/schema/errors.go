package schema

import "errors"

var (
	ErrFailure                  = errors.New(Failure.String())
	ErrInvalidToken             = errors.New(InvalidToken.String())
	ErrInvalidCaptcha           = errors.New(InvalidCaptcha.String())
	ErrInvalidAccountOrPassword = errors.New(InvalidAccountOrPassword.String())
	ErrUnauthorized             = errors.New(Unauthorized.String())
	ErrInvalidArguments         = errors.New(InvalidArguments.String())
)
