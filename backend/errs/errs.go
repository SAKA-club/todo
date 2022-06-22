package errs

import (
	"database/sql"
	"errors"
	"fmt"
)

type TodoError struct {
	ErrorCode ErrorCode `json:"code"`
	Message   string    `json:"message"`
	HttpCode  int       `json:"http_code"`
}

type ErrorCode byte

func (e ErrorCode) Error() string {
	return fmt.Sprintf("%d", e)
}

const (
	// UndefinedErr is used to signify an unknown/uninitialized status
	UndefinedErr ErrorCode = iota
	AuthErr
	ConfigErr
	DBErr
	InvalidUserErr
	NotFoundErr
	InputError
	InternalErr
)

func (e TodoError) Error() string {
	return e.Message
}

func (e TodoError) Code() string {
	return fmt.Sprintf("%d", e.ErrorCode)
}

func (e TodoError) Public() bool {
	return e.ErrorCode != DBErr && e.ErrorCode != InternalErr
}

func (e TodoError) Body() string {
	return e.Message
}

// ErrUnauthorized is a canned error for lack of authorization (http 401)
var ErrUnauthorized = TodoError{AuthErr, "invalid authorization", 401}

// ErrNotFound is a canned error for not found (http 404)
var ErrNotFound = TodoError{NotFoundErr, "not found", 404}

// ErrInvalidUser is a canned error for invalid user ID
var ErrInvalidUser = TodoError{InvalidUserErr, "invalid user", 404}

// ErrDB is a canned error for unexpected DB Errors
var ErrDB = TodoError{DBErr, "internal error", 500}

// ErrInternal is a canned error for unexpected service Errors
var ErrInternal = TodoError{InternalErr, "internal error", 500}

// ErrInvalidRequest is a canned error for invalid user ID
var ErrInvalidRequest = TodoError{InputError, "invalid request", 400}

func New(c ErrorCode, msg string, h int) error {
	return TodoError{c, msg, h}
}

func IsNotNoRows(err error) bool {
	return err != nil && !errors.Is(err, sql.ErrNoRows)
}

func IsNotFound(err error) bool {
	return err != nil && (errors.Is(err, ErrNotFound) || errors.Is(err, sql.ErrNoRows))
}

func IsUnauthorized(err error) bool {
	return err != nil && errors.Is(err, ErrUnauthorized)
}
