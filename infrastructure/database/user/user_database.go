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

func (d userDatabase) FindFollows(userID user.UserID) (user.Users, error) {
	var es []UserEntity
	err := d.db.
		Table("user_entities").
		Joins("JOIN follow_entities ON user_entities.id = follow_entities.user_id WHERE follow_entities.user_id = ?", userID).
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
	var e UserEntity
	err := d.db.Where("user_id = ?", userID).Take(&e).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, model_errors.NewInfrastructureError(err.Error())
	}
	return true, nil
}

func (d userDatabase) Delete(user user.User) error {
	e := NewUserEntity(user)
	err := d.db.Delete(&e).Error
	if err != nil {
		return model_errors.NewInfrastructureError(err.Error())
	}
	return nil
}
