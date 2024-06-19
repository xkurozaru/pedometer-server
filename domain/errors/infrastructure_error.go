package model_errors

import (
	"errors"
	"fmt"
)

var (
	ErrInfrastructure = errors.New("infrastructure error")
)

func NewInfrastructureError(message string) error {
	return fmt.Errorf("%w: %v", ErrInfrastructure, message)
}
