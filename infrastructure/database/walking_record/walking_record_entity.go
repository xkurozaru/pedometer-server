package walking_record_database

import (
	"time"

	"github.com/xkurozaru/pedometer-server/domain/common"
	"github.com/xkurozaru/pedometer-server/domain/walking_record"
	common_database "github.com/xkurozaru/pedometer-server/infrastructure/database/common"
)

type WalkingRecordEntity struct {
	common_database.Model
	UserID   string
	Date     time.Time
	Steps    int
	Distance int
	Time     int
	Calories int
}

func NewWalkingRecordEntity(w walking_record.WalkingRecord) WalkingRecordEntity {
	return WalkingRecordEntity{
		Model:    common_database.Model{ID: w.ID().String()},
		UserID:   w.UserID(),
		Date:     w.Date().Time(),
		Steps:    w.Steps(),
		Distance: w.Distance(),
		Time:     w.Time(),
		Calories: w.Calories(),
	}
}

func (e WalkingRecordEntity) ToModel() walking_record.WalkingRecord {
	return walking_record.RecreateWalkingRecord(
		common.NanoID(e.ID),
		e.UserID,
		common.DateTime(e.Date),
		e.Steps,
		e.Distance,
		e.Time,
		e.Calories,
	)
}
