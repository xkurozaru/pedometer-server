package user_database

import (
	"errors"

	model_errors "github.com/xkurozaru/pedometer-server/domain/errors"
	"github.com/xkurozaru/pedometer-server/domain/user"

	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserDatabase(db *gorm.DB) user.UserRepository {
	return userDatabase{DB: db}
}

func (d userDatabase) Create(user user.User) error {
	e := NewUserEntity(user)
	err := d.DB.Create(&e).Error
	if err != nil {
		return model_errors.NewInfrastructureError(err)
	}
	return nil
}

func (d userDatabase) Get(id string) (user.User, error) {
	var e UserEntity
	err := d.DB.Where("id = ?", id).Take(&e).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user.User{}, nil
		}
		return user.User{}, model_errors.NewInfrastructureError(err)
	}
	return e.ToModel(), nil
}

func (d userDatabase) ExistsByUserID(userID string) (bool, error) {
	var e UserEntity
	err := d.DB.Where("user_id = ?", userID).Take(&e).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, model_errors.NewInfrastructureError(err)
	}
	return true, nil
}
