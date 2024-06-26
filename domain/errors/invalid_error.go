package model_errors

import (
	"errors"
	"fmt"
)

var (
	ErrInvalid = errors.New("invalid")
)

func NewInvalidError[S ~string](message S) error {
	return fmt.Errorf("%w: %v", ErrInvalid, message)
}
