package friend_database

import (
	model_errors "github.com/xkurozaru/pedometer-server/domain/errors"
	"github.com/xkurozaru/pedometer-server/domain/friend"
	"github.com/xkurozaru/pedometer-server/domain/user"
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

func (d friendDatabase) Exists(userID, friendUserID user.UserID) (bool, error) {
	var exists bool
	err := d.db.Model(&FriendEntity{}).
		Select("COUNT(1) > 0").
		Where("user_id = ? AND friend_user_id", userID, friendUserID).
		Find(&exists).Error
	if err != nil {
		return false, model_errors.NewInfrastructureError(err.Error())
	}

	return true, nil
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
