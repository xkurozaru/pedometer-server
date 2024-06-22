package user_database

import (
	"errors"

	model_errors "github.com/xkurozaru/pedometer-server/domain/errors"
	"github.com/xkurozaru/pedometer-server/domain/user"

	"gorm.io/gorm"
)

type userDatabase struct {
	db *gorm.DB
}

func NewUserDatabase(db *gorm.DB) user.UserRepository {
	return userDatabase{db: db}
}

func (d userDatabase) Create(user user.User) error {
	e := NewUserEntity(user)
	err := d.db.Create(&e).Error
	if err != nil {
		return model_errors.NewInfrastructureError(err.Error())
	}
	return nil
}

func (d userDatabase) FindByUserID(userID user.UserID) (user.User, error) {
	var e UserEntity
	err := d.db.Where("user_id = ?", userID).Take(&e).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user.User{}, model_errors.NewNotFoundError(err.Error())
		}
		return user.User{}, model_errors.NewInfrastructureError(err.Error())
	}
	return e.ToModel(), nil
}

func (d userDatabase) FindByAuthID(authID string) (user.User, error) {
	var e UserEntity
	err := d.db.Where("auth_id = ?", authID).Take(&e).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user.User{}, model_errors.NewNotFoundError(err.Error())
		}
		return user.User{}, model_errors.NewInfrastructureError(err.Error())
	}
	return e.ToModel(), nil
}

func (d userDatabase) FindFriends(userID user.UserID) (user.Users, error) {
	var es []UserEntity
	err := d.db.Table("friend_entities").
		Select("user_entities.*").
		Joins("INNER JOIN user_entities ON friend_entities.friend_user_id = user_entities.user_id").
		Where("friend_entities.user_id = ?", userID).
		Where("friend_entities.status = ?", "established").
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

func (d userDatabase) FindAll() (user.Users, error) {
	var es []UserEntity
	err := d.db.Find(&es).Error
	if err != nil {
		return nil, model_errors.NewInfrastructureError(err.Error())
	}

	var users user.Users
	for _, e := range es {
		users = append(users, e.ToModel())
	}
	return users, nil
}

func (d userDatabase) ExistsByUserID(userID user.UserID) (bool, error) {
	var exists bool
	err := d.db.Model(&UserEntity{}).
		Select("COUNT(1) > 0").
		Where("user_id = ?", userID).
		Find(&exists).Error
	if err != nil {
		return false, model_errors.NewInfrastructureError(err.Error())
	}

	return exists, nil
}

func (d userDatabase) Delete(user user.User) error {
	e := NewUserEntity(user)
	err := d.db.Delete(&e).Error
	if err != nil {
		return model_errors.NewInfrastructureError(err.Error())
	}
	return nil
}
