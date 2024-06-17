package walking_record

import "github.com/xkurozaru/pedometer-server/domain/common"

type WalkingRecord struct {
	id       common.NanoID
	userId   string
	date     common.DateTime
	steps    int
	distance int // meters
	time     int // minutes
	calories int // kcal
}

func NewWalkingRecord(userId string, date common.DateTime, steps int, distance int, time int, calories int) WalkingRecord {
	return newWalkingRecord(common.NewNanoID(), userId, date, steps, distance, time, calories)
}
func RecreateWalkingRecord(id common.NanoID, userId string, date common.DateTime, steps int, distance int, time int, calories int) WalkingRecord {
	return newWalkingRecord(id, userId, date, steps, distance, time, calories)
}
func newWalkingRecord(id common.NanoID, userId string, date common.DateTime, steps int, distance int, time int, calories int) WalkingRecord {
	return WalkingRecord{
		id:       id,
		userId:   userId,
		date:     date,
		steps:    steps,
		distance: distance,
		time:     time,
		calories: calories,
	}
}

func (w WalkingRecord) ID() common.NanoID {
	return w.id
}
func (w WalkingRecord) UserID() string {
	return w.userId
}
func (w WalkingRecord) Date() common.DateTime {
	return w.date
}
func (w WalkingRecord) Steps() int {
	return w.steps
}
func (w WalkingRecord) Distance() int {
	return w.distance
}
func (w WalkingRecord) Time() int {
	return w.time
}
func (w WalkingRecord) Calories() int {
	return w.calories
}
