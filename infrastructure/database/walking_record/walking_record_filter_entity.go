package walking_record_database

import (
	"github.com/xkurozaru/pedometer-server/domain/common"
	"github.com/xkurozaru/pedometer-server/domain/walking_record"
	"gorm.io/gorm"
)

type WalkingRecordFilterEntity struct {
	UserIDs   []string
	StartDate string
	EndDate   string
}

func NewWalkingRecordFilterEntity(filter walking_record.WalkingRecordFilter) WalkingRecordFilterEntity {
	userIDs := []string{}
	for _, id := range filter.UserIDs {
		userIDs = append(userIDs, string(id))
	}
	return WalkingRecordFilterEntity{
		UserIDs:   userIDs,
		StartDate: filter.StartDate.Format(common.HyphenDateFormat),
		EndDate:   filter.EndDate.Format(common.HyphenDateFormat),
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
