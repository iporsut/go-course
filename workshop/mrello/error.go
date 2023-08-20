package mrello

import (
	"errors"
	"fmt"
)

type ErrorCode int

const (
	ErrCodeOther ErrorCode = iota
	ErrCodeNotFound
)

func (c ErrorCode) String() string {
	switch c {
	case ErrCodeNotFound:
		return "not found"
	case ErrCodeOther:
		return "other"
	default:
		return "unknown"
	}
}

type Error struct {
	Code    ErrorCode
	Message string
	Err     error
}

func (e *Error) Error() string {
	var unwrapStr string
	if e.Err != nil {
		unwrapStr = e.Err.Error()
	}
	return fmt.Sprintf("error code %s - message %s: %s", e.Code.String(), e.Message, unwrapStr)
}

func (e *Error) Unwrap() error {
	return e.Err
}

func WrapErr(err error, code ErrorCode, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

func IsErrCode(err error, code ErrorCode) bool {
	if err == nil {
		return false
	}

	var e *Error
	if errors.As(err, &e) {
		return e.Code == code
	}

	return false
}
