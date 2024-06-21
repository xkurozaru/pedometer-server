package follow_application

import (
	"fmt"

	"github.com/xkurozaru/pedometer-server/domain/follow"
	"github.com/xkurozaru/pedometer-server/domain/user"
)

type FollowApplicationService interface {
	Follow(userID string, followedUserID string) error
	Unfollow(userID string, followedUserID string) error
}

type followApplicationService struct {
	followRepository follow.FollowRepository
	userRepository   user.UserRepository
}

func NewFollowApplicationService(
	followRepository follow.FollowRepository,
	userRepository user.UserRepository,
) FollowApplicationService {
	return followApplicationService{
		followRepository: followRepository,
		userRepository:   userRepository,
	}
}

func (s followApplicationService) Follow(userID string, followedUserID string) error {
	follow := follow.NewFollow(userID, followedUserID)
	err := s.followRepository.Upsert(follow)
	if err != nil {
		return fmt.Errorf("Upsert: %w", err)
	}
	return nil
}

func (s followApplicationService) Unfollow(userID string, followedUserID string) error {
	follow := follow.RecreateFollow(userID, followedUserID)
	err := s.followRepository.Delete(follow)
	if err != nil {
		return fmt.Errorf("Delete: %w", err)
	}
	return nil
}
