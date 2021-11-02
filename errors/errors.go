package errors

import (
	"fmt"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("error (%d): %s", e.Code, e.Message)
}

func newError(code int, msg string) *Error {
	return &Error{Code: code, Message: msg}
}

const (
	UnknownErr = iota + 90001
)

// request error
const (
	InvalidRequestParams = iota + 10001
	InternalServer
	MissingRequestParams
	UnknownLoginType
	TokenCreateFailed
	TokenExpired
	InvalidPassword
	UserNotFound
	ActionNotAllowed
	UnAuthentication
	NotRequiredPassword
)

var (
	ErrUnknown              = newError(UnknownErr, "unknown error")
	ErrInternalServer       = newError(InternalServer, "internal server error")
	ErrInvalidRequestParams = newError(InvalidRequestParams, "invalid request params")
	ErrMissingRequestParams = newError(MissingRequestParams, "missing request params")
	ErrUnAuthentication     = newError(UnAuthentication, "unauthentication access")
	ErrTokenCreateFailed    = newError(TokenCreateFailed, "create token failed")
	ErrTokenExpired         = newError(TokenExpired, "Token is expired")
	ErrInvalidPassword      = newError(InvalidPassword, "invalid password")
	ErrUserNotFound         = newError(UserNotFound, "user not found")
	ErrActionNotAllowed     = newError(ActionNotAllowed, "action not allowed")
	ErrNotRequiredPassword  = newError(NotRequiredPassword, "passwords must be at least 6 characters long")
)
