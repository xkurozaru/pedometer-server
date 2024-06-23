package auth

import (
	"github.com/google/uuid"
	"github.com/xkurozaru/pedometer-server/domain/user"
)

type AuthRepository interface {
	Register(email string, password string) (uuid.UUID, error)
	Login(email string, password string) (string, error)
	Verify(jWT string) (uuid.UUID, error)
	Delete(u user.User) error
}
