package auth

type AuthRepository interface {
	Register(email string, password string) (string, error)
	Login(email string, password string) (string, error)
	Verify(jWT string) (string, error)
}
