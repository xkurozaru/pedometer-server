package auth

import "github.com/xkurozaru/pedometer-server/domain/user"

type AuthRepository interface {
	Register(email string, password string) (string, error)
	Login(email string, password string) (string, error)
	Verify(jWT string) (string, error)
	Delete(u user.User) error
}
