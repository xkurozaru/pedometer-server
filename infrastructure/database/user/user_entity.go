package user_database

import (
	"github.com/google/uuid"
	"github.com/xkurozaru/pedometer-server/domain/user"
)

type UserEntity struct {
	UserID   string `gorm:"primaryKey"`
	Username string
	AuthID   string `gorm:"unique"`
}

func NewUserEntity(u user.User) UserEntity {
	return UserEntity{
		UserID:   string(u.UserID()),
		Username: u.Username(),
		AuthID:   u.AuthID().String(),
	}
}

func (e UserEntity) ToModel() user.User {
	return user.RecreateUser(user.UserID(e.UserID), e.Username, uuid.MustParse(e.AuthID))
}
