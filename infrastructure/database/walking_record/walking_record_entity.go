package walking_record_database

import (
	"time"

	"github.com/xkurozaru/pedometer-server/domain/common"
	"github.com/xkurozaru/pedometer-server/domain/walking_record"
)

type WalkingRecordEntity struct {
	UserID   string `gorm:"primaryKey"`
	Date     string `gorm:"primaryKey"`
	Distance int

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func NewWalkingRecordEntity(w walking_record.WalkingRecord) WalkingRecordEntity {
	return WalkingRecordEntity{
		UserID:   w.UserID(),
		Date:     w.Date().Format(common.HyphenDateFormat),
		Distance: w.Distance(),
	}
}

func (e WalkingRecordEntity) ToModel() walking_record.WalkingRecord {
	date, _ := common.DateTimeFromString(e.Date, common.HyphenDateFormat)

	return walking_record.RecreateWalkingRecord(
		e.UserID,
		date,
		e.Distance,
	)
}
