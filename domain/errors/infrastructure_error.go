package model_errors

import (
	"errors"
	"fmt"
)

var (
	ErrInfrastructure = errors.New("infrastructure error")
)

func NewInfrastructureError(err error) error {
	return fmt.Errorf("%w: %v", ErrInfrastructure, err)
}
