package friend_application

import (
	"fmt"
	"log/slog"

	model_errors "github.com/xkurozaru/pedometer-server/domain/errors"
	"github.com/xkurozaru/pedometer-server/domain/friend"
	"github.com/xkurozaru/pedometer-server/domain/user"
)

type FriendApplicationService interface {
	RegisterFriendRequest(userID, friendUserID user.UserID) error
	AcceptFriendRequest(userID, friendUserID user.UserID) error
	FetchFriendList(userID user.UserID, status friend.FriendStatus) (user.Users, error)
	RemoveFriend(userID, friendUserID user.UserID) error
}

type friendApplicationService struct {
	friendRepository friend.FriendRepository
	friendService    friend.FriendService
	userRepository   user.UserRepository
}

func NewFriendApplicationService(
	friendRepository friend.FriendRepository,
	friendService friend.FriendService,
	userRepository user.UserRepository,
) FriendApplicationService {
	return friendApplicationService{
		friendRepository: friendRepository,
		friendService:    friendService,
		userRepository:   userRepository,
	}
}

func (s friendApplicationService) RegisterFriendRequest(
	userID user.UserID,
	friendUserID user.UserID,
) error {
	exists, err := s.userRepository.ExistsByUserID(friendUserID)
	if err != nil {
		return fmt.Errorf("ExistsByUserID: %w", err)
	}
	slog.Info("ExistsByUserID", "exists", exists)

	if !exists {
		return model_errors.NewNotFoundError(string(friendUserID))
	}

	exists, err = s.friendRepository.Exists(userID, friendUserID)
	if err != nil {
		return fmt.Errorf("Exists: %w", err)
	}
	slog.Info("Exists", "exists", exists)

	if exists {
		err = s.friendService.EstablishPairIfRequested(userID, friendUserID)
		if err != nil {
			return fmt.Errorf("EstablishIfRequested: %w", err)
		}
		slog.Info("EstablishIfRequested")

		return nil
	}

	friends := friend.NewFriendPair(userID, friendUserID)
	slog.Info("NewFriendPair", "friends", friends)

	err = s.friendRepository.UpsertAll(friends)
	if err != nil {
		return fmt.Errorf("UpsertAll: %w", err)
	}
	slog.Info("UpsertAll")

	return nil
}

func (s friendApplicationService) AcceptFriendRequest(
	userID user.UserID,
	friendUserID user.UserID,
) error {
	exists, err := s.friendRepository.Exists(userID, friendUserID)
	if err != nil {
		return fmt.Errorf("Exists: %w", err)
	}
	slog.Info("Exists", "exists", exists)

	if !exists {
		return model_errors.NewNotFoundError(string(friendUserID))
	}

	err = s.friendService.EstablishPairIfRequested(userID, friendUserID)
	if err != nil {
		return fmt.Errorf("EstablishIfRequested: %w", err)
	}
	slog.Info("EstablishIfRequested")

	return nil
}

func (s friendApplicationService) FetchFriendList(
	userID user.UserID,
	status friend.FriendStatus,
) (user.Users, error) {
	friends, err := s.friendRepository.FindFriendUsers(userID, status)
	if err != nil {
		return nil, fmt.Errorf("FindFriends: %w", err)
	}
	slog.Info("FindFriends", "friends", friends)

	return friends, nil
}

func (s friendApplicationService) RemoveFriend(
	userID user.UserID,
	friendUserID user.UserID,
) error {
	exists, err := s.friendRepository.Exists(userID, friendUserID)
	if err != nil {
		return fmt.Errorf("Exists: %w", err)
	}
	slog.Info("Exists", "exists", exists)

	if !exists {
		return model_errors.NewNotFoundError(string(friendUserID))
	}

	err = s.friendService.DeletePair(userID, friendUserID)
	if err != nil {
		return fmt.Errorf("DeletePair: %w", err)
	}
	slog.Info("DeletePair")

	return nil
}
