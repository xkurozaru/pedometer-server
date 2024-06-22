package friend_database

import (
	friend_application "github.com/xkurozaru/pedometer-server/application/friend"
	model_errors "github.com/xkurozaru/pedometer-server/domain/errors"
	"github.com/xkurozaru/pedometer-server/domain/user"
	"gorm.io/gorm"
)

type friendQueryService struct {
	db *gorm.DB
}

func NewFriendQueryService(db *gorm.DB) friend_application.FriendQueryService {
	return friendQueryService{db: db}
}

func (s friendQueryService) GetFriendList(userID user.UserID) (friend_application.FriendListDTO, error) {
	type Entity struct {
		UserID         string
		FriendUserID   string
		FriendUsername string
		Status         string
	}

	var e []Entity

	err := s.db.Table("friend_entities").
		Select("friend_entities.user_id, friend_entities.friend_user_id, user_entities.username as friend_username, friend_entities.status").
		Joins("JOIN user_entities ON friend_entities.friend_user_id = user_entities.user_id").
		Where("friend_entities.user_id = ?", userID).
		Find(&e).Error
	if err != nil {
		return friend_application.FriendListDTO{}, model_errors.NewInfrastructureError(err.Error())
	}

	var friends []friend_application.FriendDTO
	var requested []friend_application.FriendDTO
	var requesting []friend_application.FriendDTO

	for _, entity := range e {
		friend := friend_application.FriendDTO{
			FriendUserID:   entity.FriendUserID,
			FriendUsername: entity.FriendUsername,
		}

		switch entity.Status {
		case "established":
			friends = append(friends, friend)
		case "requested":
			requested = append(requested, friend)
		case "requesting":
			requesting = append(requesting, friend)
		}
	}

	return friend_application.FriendListDTO{
		Friends:    friends,
		Requested:  requested,
		Requesting: requesting,
	}, nil
}
