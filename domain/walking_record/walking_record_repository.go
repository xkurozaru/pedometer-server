package walking_record

type WalkingRecordRepository interface {
	Upsert(w WalkingRecord) error
	AllUpsert(ws WalkingRecords) error
	FindByFilter(filter WalkingRecordFilter, order WalkingRecordOrder) (WalkingRecords, error)
}
