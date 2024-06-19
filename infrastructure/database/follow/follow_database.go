package follow_database

import (
	model_errors "github.com/xkurozaru/pedometer-server/domain/errors"
	"github.com/xkurozaru/pedometer-server/domain/follow"
	"gorm.io/gorm"
)

type followDatabase struct {
	db *gorm.DB
}

func NewFollowDatabase(db *gorm.DB) follow.FollowRepository {
	return followDatabase{db: db}
}

func (d followDatabase) Upsert(follow follow.Follow) error {
	e := NewFollowEntity(follow)
	err := d.db.Save(&e).Error
	if err != nil {
		return model_errors.NewInfrastructureError(err.Error())
	}
	return nil
}

func (d followDatabase) FindByUserID(userID string) (follow.Follows, error) {
	var es []FollowEntity
	err := d.db.Where("user_id = ?", userID).Find(&es).Error
	if err != nil {
		return nil, model_errors.NewInfrastructureError(err.Error())
	}

	var fs follow.Follows
	for _, e := range es {
		fs = append(fs, e.ToModel())
	}
	return fs, nil
}

func (d followDatabase) Delete(follow follow.Follow) error {
	e := NewFollowEntity(follow)
	err := d.db.Delete(&e).Error
	if err != nil {
		return model_errors.NewInfrastructureError(err.Error())
	}
	return nil
}
