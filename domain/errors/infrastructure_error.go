package model_errors

import (
	"errors"
	"fmt"
)

var (
	ErrInfrastructure = errors.New("infrastructure error")
)

func NewInfrastructureError[S ~string](message S) error {
	return fmt.Errorf("%w: %v", ErrInfrastructure, message)
}
