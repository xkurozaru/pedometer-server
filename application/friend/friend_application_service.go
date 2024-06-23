package friend_application

import (
	"fmt"

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
}

func NewFriendApplicationService(
	friendRepository friend.FriendRepository,
	friendService friend.FriendService,
) FriendApplicationService {
	return friendApplicationService{
		friendRepository: friendRepository,
		friendService:    friendService,
	}
}

func (s friendApplicationService) RegisterFriendRequest(
	userID user.UserID,
	friendUserID user.UserID,
) error {
	exists, err := s.friendRepository.Exists(userID, friendUserID)
	if err != nil {
		return fmt.Errorf("Exists: %w", err)
	}
	if exists {
		err = s.friendService.EstablishPairIfRequested(userID, friendUserID)
		if err != nil {
			return fmt.Errorf("EstablishIfRequested: %w", err)
		}
		return nil
	}

	friends := friend.NewFriendPair(userID, friendUserID)
	err = s.friendRepository.UpsertAll(friends)
	if err != nil {
		return fmt.Errorf("UpsertAll: %w", err)
	}

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
	if !exists {
		return model_errors.NewNotFoundError(string(friendUserID))
	}

	err = s.friendService.EstablishPairIfRequested(userID, friendUserID)
	if err != nil {
		return fmt.Errorf("EstablishIfRequested: %w", err)
	}

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
	if !exists {
		return model_errors.NewNotFoundError(string(friendUserID))
	}

	err = s.friendService.DeletePair(userID, friendUserID)
	if err != nil {
		return fmt.Errorf("DeletePair: %w", err)
	}

	return nil
}
