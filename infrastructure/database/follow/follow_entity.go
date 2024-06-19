package follow_database

import (
	"time"

	"github.com/xkurozaru/pedometer-server/domain/follow"
)

type FollowEntity struct {
	UserID         string `gorm:"primaryKey"`
	FollowedUserID string `gorm:"primaryKey"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func NewFollowEntity(f follow.Follow) FollowEntity {
	return FollowEntity{
		UserID:         f.UserID(),
		FollowedUserID: f.FollowedUserID(),
	}
}

func (e FollowEntity) ToModel() follow.Follow {
	return follow.RecreateFollow(e.UserID, e.FollowedUserID)
}
