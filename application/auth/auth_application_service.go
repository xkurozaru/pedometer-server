package auth_application

import "github.com/xkurozaru/pedometer-server/domain/auth"

type AuthApplicationService interface {
	Login(email string, password string) (string, error)
}

type authApplicationService struct {
	authRepository auth.AuthRepository
}

func NewAuthApplicationService(
	authRepository auth.AuthRepository,
) AuthApplicationService {
	return authApplicationService{
		authRepository: authRepository,
	}
}

func (s authApplicationService) Login(email string, password string) (string, error) {
	token, err := s.authRepository.Login(email, password)
	if err != nil {
		return "", err
	}

	return token, nil
}
