package user_application

import (
	"fmt"

	"github.com/xkurozaru/pedometer-server/domain/auth"
	model_errors "github.com/xkurozaru/pedometer-server/domain/errors"
	"github.com/xkurozaru/pedometer-server/domain/user"
)

type UserApplicationService interface {
	RegisterUser(email string, password string, userID user.UserID, username string) error
	FetchUserByToken(token string) (user.User, error)
	FetchUserByUserID(userID user.UserID) (user.User, error)
	Delete(u user.User) error
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
	userID user.UserID,
	username string,
) error {
	exists, err := s.userRepository.ExistsByUserID(userID)
	if err != nil {
		return fmt.Errorf("ExistsByUserID: %w", err)
	}
	if exists {
		return model_errors.NewAlreadyExistsError(string(userID))
	}

	authID, err := s.authRepository.Register(email, password)
	if err != nil {
		return fmt.Errorf("Register: %w", err)
	}

	u := user.NewUser(userID, username, authID)

	err = s.userRepository.Create(u)
	if err != nil {
		return fmt.Errorf("Create: %w", err)
	}

	return nil
}

func (s userApplicationService) FetchUserByToken(token string) (user.User, error) {
	authID, err := s.authRepository.Verify(token)
	if err != nil {
		return user.User{}, fmt.Errorf("Verify: %w", err)
	}

	u, err := s.userRepository.FindByAuthID(authID)
	if err != nil {
		return user.User{}, fmt.Errorf("Get: %w", err)
	}
	return u, nil
}

func (s userApplicationService) FetchUserByUserID(userID user.UserID) (user.User, error) {
	u, err := s.userRepository.FindByUserID(userID)
	if err != nil {
		return user.User{}, fmt.Errorf("FindByUserID: %w", err)
	}
	return u, nil
}

func (s userApplicationService) Delete(u user.User) error {
	err := s.authRepository.Delete(u)
	if err != nil {
		return fmt.Errorf("DeleteByAuthID: %w", err)
	}

	err = s.userRepository.Delete(u)
	if err != nil {
		return fmt.Errorf("Delete: %w", err)
	}

	return nil
}
