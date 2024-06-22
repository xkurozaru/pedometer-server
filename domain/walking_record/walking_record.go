package walking_record

import (
	"github.com/xkurozaru/pedometer-server/domain/common"
	"github.com/xkurozaru/pedometer-server/domain/user"
)

type WalkingRecord struct {
	userId   user.UserID
	date     common.DateTime
	distance int // meters
}

func NewWalkingRecord(userId user.UserID, date common.DateTime, distance int) WalkingRecord {
	return newWalkingRecord(userId, date.StartOfDay(), distance)
}
func RecreateWalkingRecord(userId user.UserID, date common.DateTime, distance int) WalkingRecord {
	return newWalkingRecord(userId, date, distance)
}
func newWalkingRecord(userId user.UserID, date common.DateTime, distance int) WalkingRecord {
	return WalkingRecord{
		userId:   userId,
		date:     date,
		distance: distance,
	}
}

func (w WalkingRecord) UserID() user.UserID {
	return w.userId
}
func (w WalkingRecord) Date() common.DateTime {
	return w.date
}
func (w WalkingRecord) Distance() int {
	return w.distance
}

type WalkingRecords []WalkingRecord

func (ws WalkingRecords) TotalUserDistance() map[user.UserID]int {
	totals := map[user.UserID]int{}
	for _, w := range ws {
		totals[w.UserID()] += w.Distance()
	}
	return totals
}
