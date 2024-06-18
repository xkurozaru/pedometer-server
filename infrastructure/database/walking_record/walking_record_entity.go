package walking_record_database

import (
	"time"

	"github.com/xkurozaru/pedometer-server/domain/common"
	"github.com/xkurozaru/pedometer-server/domain/walking_record"
)

type WalkingRecordEntity struct {
	UserID   string
	Date     time.Time
	Distance int
}

func NewWalkingRecordEntity(w walking_record.WalkingRecord) WalkingRecordEntity {
	return WalkingRecordEntity{
		UserID:   w.UserID(),
		Date:     w.Date().Time(),
		Distance: w.Distance(),
	}
}

func (e WalkingRecordEntity) ToModel() walking_record.WalkingRecord {
	return walking_record.RecreateWalkingRecord(
		e.UserID,
		common.DateTime(e.Date),
		e.Distance,
	)
}
