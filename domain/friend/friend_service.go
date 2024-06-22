package friend

import (
	"fmt"

	model_errors "github.com/xkurozaru/pedometer-server/domain/errors"
	"github.com/xkurozaru/pedometer-server/domain/user"
)

type FriendService interface {
	EstablishIfRequested(fromUserID, toUserID user.UserID) error
	DeletePair(fromUserID, toUserID user.UserID) error
}

type friendService struct {
	friendRepository FriendRepository
}

func NewFriendService(
	friendRepository FriendRepository,
) FriendService {
	return friendService{
		friendRepository: friendRepository,
	}
}

func (s friendService) EstablishIfRequested(
	fromUserID user.UserID,
	toUserID user.UserID,
) error {
	friend, err := s.friendRepository.Find(toUserID, fromUserID)
	if err != nil {
		return fmt.Errorf("Find: %w", err)
	}
	if !friend.IsRequested() {
		return model_errors.NewAlreadyExistsError(string(toUserID))
	}

	friends := friend.Establish()
	err = s.friendRepository.UpsertAll(friends)
	if err != nil {
		return fmt.Errorf("UpsertAll: %w", err)
	}

	return nil
}

func (s friendService) DeletePair(
	fromUserID user.UserID,
	toUserID user.UserID,
) error {
	err := s.friendRepository.Delete(fromUserID, toUserID)
	if err != nil {
		return fmt.Errorf("Delete: %w", err)
	}

	err = s.friendRepository.Delete(toUserID, fromUserID)
	if err != nil {
		return fmt.Errorf("Delete: %w", err)
	}

	return nil
}
