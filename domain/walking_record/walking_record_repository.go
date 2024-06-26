package walking_record

type WalkingRecordRepository interface {
	UpsertAll(ws WalkingRecords) error
	FindByFilter(filter WalkingRecordFilter, order WalkingRecordOrder) (WalkingRecords, error)
}
