package model_errors

import (
	"errors"
	"fmt"
)

var (
	ErrAlreadyExists = errors.New("already exists")
)

func NewAlreadyExistsError[S ~string](message S) error {
	return fmt.Errorf("%w: %v", ErrAlreadyExists, message)
}
