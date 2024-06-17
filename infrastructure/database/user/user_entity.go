package user_database

import (
	"github.com/xkurozaru/pedometer-server/domain/user"
	common_database "github.com/xkurozaru/pedometer-server/infrastructure/database/common"
)

type UserEntity struct {
	common_database.Model
	UserID   string
	Username string
}

func NewUserEntity(u user.User) UserEntity {
	return UserEntity{
		Model:    common_database.Model{ID: u.ID()},
		UserID:   u.UserID(),
		Username: u.Username(),
	}
}

func (e UserEntity) ToModel() user.User {
	return user.RecreateUser(e.ID, e.UserID, e.Username)
}
