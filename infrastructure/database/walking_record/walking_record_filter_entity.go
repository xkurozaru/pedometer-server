package walking_record_database

import (
	"time"

	"github.com/xkurozaru/pedometer-server/domain/walking_record"
	"gorm.io/gorm"
)

type WalkingRecordFilterEntity struct {
	UserIDs   []string
	StartDate time.Time
	EndDate   time.Time
}

func NewWalkingRecordFilterEntity(filter walking_record.WalkingRecordFilter) WalkingRecordFilterEntity {
	return WalkingRecordFilterEntity{
		UserIDs:   filter.UserIDs,
		StartDate: filter.StartDate.Time(),
		EndDate:   filter.EndDate.Time(),
	}
}

func (e WalkingRecordFilterEntity) FilterQuery(db *gorm.DB) *gorm.DB {
	return db.Scopes(
		e.whereUserIDs,
		e.whereBetweenDate,
	)
}

func (e WalkingRecordFilterEntity) whereUserIDs(db *gorm.DB) *gorm.DB {
	if len(e.UserIDs) == 0 {
		return db
	}
	return db.Where("user_id IN (?)", e.UserIDs)
}

func (e WalkingRecordFilterEntity) whereBetweenDate(db *gorm.DB) *gorm.DB {
	return db.Where("date BETWEEN ? AND ?", e.StartDate, e.EndDate)
}
