package walking_record_database

import (
	model_errors "github.com/xkurozaru/pedometer-server/domain/errors"
	"github.com/xkurozaru/pedometer-server/domain/walking_record"
	"gorm.io/gorm"
)

type walkingRecordDatabase struct {
	DB *gorm.DB
}

func NewWalkingRecordDatabase(db *gorm.DB) walking_record.WalkingRecordRepository {
	return walkingRecordDatabase{DB: db}
}

func (d walkingRecordDatabase) Upsert(w walking_record.WalkingRecord) error {
	e := NewWalkingRecordEntity(w)
	err := d.DB.Save(&e).Error
	if err != nil {
		return model_errors.NewInfrastructureError(err)
	}
	return nil
}
