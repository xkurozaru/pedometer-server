package model_errors

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound = errors.New("not found")
)

func NewNotFoundError[S ~string](message S) error {
	return fmt.Errorf("%w: %v", ErrNotFound, message)
}
