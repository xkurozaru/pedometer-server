package walking_record

type WalkingRecordRepository interface {
	Upsert(w WalkingRecord) error
}
