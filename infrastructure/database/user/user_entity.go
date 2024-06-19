package user_database

import (
	"github.com/xkurozaru/pedometer-server/domain/user"
)

type UserEntity struct {
	ID       string `gorm:"primaryKey"`
	UserID   string
	Username string
}

func NewUserEntity(u user.User) UserEntity {
	return UserEntity{
		ID:       u.ID(),
		UserID:   u.UserID(),
		Username: u.Username(),
	}
}

func (e UserEntity) ToModel() user.User {
	return user.RecreateUser(e.ID, e.UserID, e.Username)
}
