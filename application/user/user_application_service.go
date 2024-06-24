package user_application

import (
	"fmt"
	"log/slog"

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
	slog.Info("ExistByUserID", "exists", exists)

	if exists {
		return model_errors.NewAlreadyExistsError(string(userID))
	}

	authID, err := s.authRepository.Register(email, password)
	if err != nil {
		return fmt.Errorf("Register: %w", err)
	}
	slog.Info("Register", "authID", authID)

	u := user.NewUser(userID, username, authID)
	slog.Info("NewUser", "user", u)

	err = s.userRepository.Create(u)
	if err != nil {
		return fmt.Errorf("Create: %w", err)
	}
	slog.Info("Create")

	return nil
}

func (s userApplicationService) FetchUserByToken(token string) (user.User, error) {
	authID, err := s.authRepository.Verify(token)
	if err != nil {
		return user.User{}, fmt.Errorf("Verify: %w", err)
	}
	slog.Info("Verify", "authID", authID)

	u, err := s.userRepository.FindByAuthID(authID)
	if err != nil {
		return user.User{}, fmt.Errorf("Get: %w", err)
	}
	slog.Info("FindByAuthID", "user", u)

	return u, nil
}

func (s userApplicationService) FetchUserByUserID(userID user.UserID) (user.User, error) {
	u, err := s.userRepository.FindByUserID(userID)
	if err != nil {
		return user.User{}, fmt.Errorf("FindByUserID: %w", err)
	}
	slog.Info("FindByUserID", "user", u)

	return u, nil
}

func (s userApplicationService) Delete(u user.User) error {
	err := s.authRepository.Delete(u)
	if err != nil {
		return fmt.Errorf("DeleteByAuthID: %w", err)
	}
	slog.Info("DeleteByAuthID")

	err = s.userRepository.Delete(u)
	if err != nil {
		return fmt.Errorf("Delete: %w", err)
	}
	slog.Info("Delete")

	return nil
}
