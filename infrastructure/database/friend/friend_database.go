package friend_database

import (
	model_errors "github.com/xkurozaru/pedometer-server/domain/errors"
	"github.com/xkurozaru/pedometer-server/domain/friend"
	"github.com/xkurozaru/pedometer-server/domain/user"
	user_database "github.com/xkurozaru/pedometer-server/infrastructure/database/user"
	"gorm.io/gorm"
)

type friendDatabase struct {
	db *gorm.DB
}

func NewFriendDatabase(db *gorm.DB) friend.FriendRepository {
	return friendDatabase{db: db}
}

func (d friendDatabase) Find(userID, friendUserID user.UserID) (friend.Friend, error) {
	var e FriendEntity
	err := d.db.Where("user_id = ? AND friend_user_id = ?", userID, friendUserID).Take(&e).Error
	if err != nil {
		return friend.Friend{}, model_errors.NewInfrastructureError(err.Error())
	}

	return e.ToModel(), nil
}

func (d friendDatabase) FindFriends(userID user.UserID, status friend.FriendStatus) (user.Users, error) {
	var es []user_database.UserEntity
	err := d.db.Table("friend_entities").
		Select("user_entities.*").
		Joins("INNER JOIN user_entities ON friend_entities.friend_user_id = user_entities.user_id").
		Where("friend_entities.user_id = ?", userID).
		Where("friend_entities.status = ?", status.ToString()).
		Find(&es).Error
	if err != nil {
		return nil, model_errors.NewInfrastructureError(err.Error())
	}

	var users user.Users
	for _, e := range es {
		users = append(users, e.ToModel())
	}
	return users, nil
}

func (d friendDatabase) Exists(userID, friendUserID user.UserID) (bool, error) {
	var exists bool
	err := d.db.Model(&FriendEntity{}).
		Select("COUNT(1) > 0").
		Where("user_id = ? AND friend_user_id = ?", userID, friendUserID).
		Limit(1).
		Find(&exists).Error
	if err != nil {
		return false, model_errors.NewInfrastructureError(err.Error())
	}

	return exists, nil
}

func (d friendDatabase) UpsertAll(friends friend.Friends) error {
	var es []FriendEntity
	for _, f := range friends {
		es = append(es, NewFriendEntity(f))
	}

	err := d.db.Save(&es).Error
	if err != nil {
		return model_errors.NewInfrastructureError(err.Error())
	}

	return nil
}

func (d friendDatabase) Delete(userID, friendUserID user.UserID) error {
	err := d.db.Where("user_id = ? AND friend_user_id = ?", userID, friendUserID).Delete(&FriendEntity{}).Error
	if err != nil {
		return model_errors.NewInfrastructureError(err.Error())
	}

	return nil
}
