package friend_application

import (
	"fmt"

	model_errors "github.com/xkurozaru/pedometer-server/domain/errors"
	"github.com/xkurozaru/pedometer-server/domain/friend"
	"github.com/xkurozaru/pedometer-server/domain/user"
)

type FriendApplicationService interface {
	RegisterFriendRequest(fromUserID user.UserID, toUserID user.UserID) error
	AcceptFriendRequest(fromUserID user.UserID, toUserID user.UserID) error
	FetchFriendList(userID user.UserID) (FriendListDTO, error)
	RemoveFriend(fromUserID user.UserID, toUserID user.UserID) error
}

type friendApplicationService struct {
	friendRepository   friend.FriendRepository
	friendService      friend.FriendService
	friendQueryService FriendQueryService
}

func NewFriendApplicationService(
	friendRepository friend.FriendRepository,
	friendService friend.FriendService,
	friendQueryService FriendQueryService,
) FriendApplicationService {
	return friendApplicationService{
		friendRepository:   friendRepository,
		friendService:      friendService,
		friendQueryService: friendQueryService,
	}
}

func (s friendApplicationService) RegisterFriendRequest(
	fromUserID user.UserID,
	toUserID user.UserID,
) error {
	exists, err := s.friendRepository.Exists(fromUserID, toUserID)
	if err != nil {
		return fmt.Errorf("Exists: %w", err)
	}
	if exists {
		err = s.friendService.EstablishIfRequested(fromUserID, toUserID)
		if err != nil {
			return fmt.Errorf("EstablishIfRequested: %w", err)
		}
		return nil
	}

	friends := friend.NewFriendPair(fromUserID, toUserID)
	err = s.friendRepository.UpsertAll(friends)
	if err != nil {
		return fmt.Errorf("UpsertAll: %w", err)
	}

	return nil
}

func (s friendApplicationService) AcceptFriendRequest(
	fromUserID user.UserID,
	toUserID user.UserID,
) error {
	exists, err := s.friendRepository.Exists(fromUserID, toUserID)
	if err != nil {
		return fmt.Errorf("Exists: %w", err)
	}
	if !exists {
		return model_errors.NewNotFoundError(string(fromUserID))
	}

	err = s.friendService.EstablishIfRequested(fromUserID, toUserID)
	if err != nil {
		return fmt.Errorf("EstablishIfRequested: %w", err)
	}

	return nil
}

func (s friendApplicationService) FetchFriendList(
	userID user.UserID,
) (FriendListDTO, error) {
	list, err := s.friendQueryService.GetFriendList(userID)
	if err != nil {
		return FriendListDTO{}, fmt.Errorf("GetFriendList: %w", err)
	}

	return list, nil
}

func (s friendApplicationService) RemoveFriend(
	fromUserID user.UserID,
	toUserID user.UserID,
) error {
	exists, err := s.friendRepository.Exists(fromUserID, toUserID)
	if err != nil {
		return fmt.Errorf("Exists: %w", err)
	}
	if !exists {
		return model_errors.NewNotFoundError(string(fromUserID))
	}

	err = s.friendService.DeletePair(fromUserID, toUserID)
	if err != nil {
		return fmt.Errorf("DeletePair: %w", err)
	}

	return nil
}
