package model_errors

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound = errors.New("not found")
)

func NewNotFoundError(message string) error {
	return fmt.Errorf("%w: %v", ErrNotFound, message)
}
