package errs

import (
	"database/sql"
	"errors"
	"fmt"
)

type TodoError struct {
	ErrorCode ErrorCode `json:"code"`
	Message   string    `json:"message"`
}

type ErrorCode byte

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
var ErrUnauthorized = TodoError{AuthErr, "invalid authorization"}

// ErrNotFound is a canned error for not found (http 404)
var ErrNotFound = TodoError{NotFoundErr, "not found"}

// ErrInvalidUser is a canned error for invalid user ID
var ErrInvalidUser = TodoError{InvalidUserErr, "invalid user"}

// ErrDB is a canned error for unexpected DB Errors
var ErrDB = TodoError{DBErr, "internal error"}

// ErrInternal is a canned error for unexpected service Errors
var ErrInternal = TodoError{InternalErr, "internal error"}

// ErrInvalidRequest is a canned error for invalid user ID
var ErrInvalidRequest = TodoError{InputError, "invalid request"}

func New(c ErrorCode, msg string) error {
	return TodoError{c, msg}
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
