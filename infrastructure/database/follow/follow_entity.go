package follow_database

import (
	"time"

	"github.com/xkurozaru/pedometer-server/domain/follow"
	"github.com/xkurozaru/pedometer-server/domain/user"
)

type FollowEntity struct {
	UserID         string `gorm:"primaryKey"`
	FollowedUserID string `gorm:"primaryKey"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func NewFollowEntity(f follow.Follow) FollowEntity {
	return FollowEntity{
		UserID:         string(f.UserID()),
		FollowedUserID: string(f.FollowedUserID()),
	}
}

func (e FollowEntity) ToModel() follow.Follow {
	return follow.RecreateFollow(user.UserID(e.UserID), user.UserID(e.FollowedUserID))
}
