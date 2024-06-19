package model_errors

import (
	"errors"
	"fmt"
)

var (
	ErrInvalid = errors.New("invalid")
)

func NewInvalidError(message string) error {
	return fmt.Errorf("%w: %v", ErrInvalid, message)
}
