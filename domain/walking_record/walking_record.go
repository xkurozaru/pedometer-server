package walking_record

import "github.com/xkurozaru/pedometer-server/domain/common"

type WalkingRecord struct {
	userId   string
	date     common.DateTime
	distance int // meters
}

func NewWalkingRecord(userId string, date common.DateTime, distance int) WalkingRecord {
	return newWalkingRecord(userId, date, distance)
}
func RecreateWalkingRecord(userId string, date common.DateTime, distance int) WalkingRecord {
	return newWalkingRecord(userId, date, distance)
}
func newWalkingRecord(userId string, date common.DateTime, distance int) WalkingRecord {
	return WalkingRecord{
		userId:   userId,
		date:     date,
		distance: distance,
	}
}

func (w WalkingRecord) UserID() string {
	return w.userId
}
func (w WalkingRecord) Date() common.DateTime {
	return w.date
}
func (w WalkingRecord) Distance() int {
	return w.distance
}
