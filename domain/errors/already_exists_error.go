package model_errors

import (
	"errors"
	"fmt"
)

var (
	ErrAlreadyExists = errors.New("already exists")
)

func NewAlreadyExistsError(message string) error {
	return fmt.Errorf("%w: %v", ErrAlreadyExists, message)
}
