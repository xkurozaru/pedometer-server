package walking_record

type WalkingRecordOrder int

const (
	_ WalkingRecordOrder = iota
	WalkingRecordOrderUserIDAsc
	WalkingRecordOrderUserIDDesc
	WalkingRecordOrderDateAsc
	WalkingRecordOrderDateDesc
	WalkingRecordOrderDistanceAsc
	WalkingRecordOrderDistanceDesc
)
