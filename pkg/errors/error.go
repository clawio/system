package errors

import (
	"fmt"
)

type ErrorCode int

const (
	InternalErrorCode ErrorCode = 1 + iota
	UnauthorizedErrorCode
	NotFoundErrorCode
	IsDirectoryErrorCode
)

var codes = [...]string{
	"internal server error",
	"unauthorized",
	"not found",
	"is directory",
}

func (e ErrorCode) String() string {
	return codes[e-1]
}

type Error struct {
	Code    ErrorCode
	Message string
}

func (e *Error) String() string {
	if e.Message == "" {
		return fmt.Sprintf("%d: %s", e.Code, e.Code.String())
	}
	return fmt.Sprintf("%d: %s(%s)", e.Code, e.Code.String(), e.Message)
}

func (e *Error) Error() string {
	return e.String()
}

func NewError(code ErrorCode, msg string) *Error {
	return &Error{Code: code, Message: msg}
}
