package application

import (
	"fmt"

	"github.com/HackU-team04/Pedometer/domain/auth"
	model_errors "github.com/HackU-team04/Pedometer/domain/errors"
	"github.com/HackU-team04/Pedometer/domain/user"
)

type UserApplicationService interface {
	RegisterUser(email string, password string, userID string, username string) error
	GetUser(token string) (user.User, error)
}

type userApplicationService struct {
	authRepository auth.AuthRepository
	userRepository user.UserRepository
}

func NewUserApplicationService(
	authRepository auth.AuthRepository,
	userRepository user.UserRepository,
) UserApplicationService {
	return userApplicationService{
		authRepository: authRepository,
		userRepository: userRepository,
	}
}

func (s userApplicationService) RegisterUser(
	email string,
	password string,
	userID string,
	username string,
) error {
	exists, err := s.userRepository.ExistsByUserID(userID)
	if err != nil {
		return fmt.Errorf("ExistsByUserID: %w", err)
	}
	if exists {
		return model_errors.NewAlreadyExistsError("user already exists")
	}

	id, err := s.authRepository.Register(email, password)
	if err != nil {
		return fmt.Errorf("Register: %w", err)
	}

	u := user.NewUser(id, userID, username)

	err = s.userRepository.Create(u)
	if err != nil {
		return fmt.Errorf("Create: %w", err)
	}

	return nil
}

func (s userApplicationService) GetUser(token string) (user.User, error) {
	id, err := s.authRepository.Verify(token)
	if err != nil {
		return user.User{}, fmt.Errorf("Verify: %w", err)
	}

	u, err := s.userRepository.Get(id)
	if err != nil {
		return user.User{}, fmt.Errorf("Get: %w", err)
	}
	return u, nil
}
