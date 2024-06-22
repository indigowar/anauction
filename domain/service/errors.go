package service

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidCredentials = errors.New("provided credentials are invalid")
	ErrInvalidData        = errors.New("provided data is invalid")
	ErrInternal           = errors.New("internal server error")
)

type DuplicationError struct {
	Object string
	Field  string
}

var _ error = &DuplicationError{}

// Error implements error.
func (e *DuplicationError) Error() string {
	return fmt.Sprintf("%s already exists in the storage with provided %s", e.Object, e.Field)
}

func NewDuplicationError(object string, field string) *DuplicationError {
	return &DuplicationError{
		Object: object,
		Field:  field,
	}
}
