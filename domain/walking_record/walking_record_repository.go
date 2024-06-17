package walking_record

type WalkingRecordRepository interface {
	Create(w WalkingRecord) error
}
