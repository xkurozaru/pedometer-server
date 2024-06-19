package walking_record_database

import (
	"github.com/xkurozaru/pedometer-server/domain/walking_record"
	"gorm.io/gorm"
)

type WalkingRecordOrderEntity int

const (
	_ WalkingRecordOrderEntity = iota
	WalkingRecordOrderEntityUserIDAsc
	WalkingRecordOrderEntityUserIDDesc
	WalkingRecordOrderEntityDateAsc
	WalkingRecordOrderEntityDateDesc
	WalkingRecordOrderEntityDistanceAsc
	WalkingRecordOrderEntityDistanceDesc
)

func NewWalkingRecordOrderEntity(order walking_record.WalkingRecordOrder) WalkingRecordOrderEntity {
	return WalkingRecordOrderEntity(order)
}

func (w WalkingRecordOrderEntity) OrderQuery(db *gorm.DB) *gorm.DB {
	switch w {
	case WalkingRecordOrderEntityUserIDAsc:
		return db.Order("user_id ASC")
	case WalkingRecordOrderEntityUserIDDesc:
		return db.Order("user_id DESC")
	case WalkingRecordOrderEntityDateAsc:
		return db.Order("date ASC")
	case WalkingRecordOrderEntityDateDesc:
		return db.Order("date DESC")
	case WalkingRecordOrderEntityDistanceAsc:
		return db.Order("distance ASC")
	case WalkingRecordOrderEntityDistanceDesc:
		return db.Order("distance DESC")
	default:
		return db
	}

}
