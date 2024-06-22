package friend_database

import (
	"time"

	"github.com/xkurozaru/pedometer-server/domain/friend"
	"github.com/xkurozaru/pedometer-server/domain/user"
)

type FriendEntity struct {
	UserID       string `gorm:"primaryKey"`
	FriendUserID string `gorm:"primaryKey"`
	Status       string

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func NewFriendEntity(f friend.Friend) FriendEntity {
	return FriendEntity{
		UserID:       string(f.UserID()),
		FriendUserID: string(f.FriendUserID()),
		Status:       f.Status().ToString(),
	}
}

func (e FriendEntity) ToModel() friend.Friend {
	return friend.RecreateFriend(
		user.UserID(e.UserID),
		user.UserID(e.FriendUserID),
		friend.FriendStatusFromString(e.Status),
	)
}
