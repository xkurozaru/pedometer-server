package walking_record_database

import (
	model_errors "github.com/xkurozaru/pedometer-server/domain/errors"
	"github.com/xkurozaru/pedometer-server/domain/walking_record"
	"gorm.io/gorm"
)

type walkingRecordDatabase struct {
	db *gorm.DB
}

func NewWalkingRecordDatabase(db *gorm.DB) walking_record.WalkingRecordRepository {
	return walkingRecordDatabase{db: db}
}

func (d walkingRecordDatabase) Upsert(w walking_record.WalkingRecord) error {
	e := NewWalkingRecordEntity(w)
	err := d.db.Save(&e).Error
	if err != nil {
		return model_errors.NewInfrastructureError(err.Error())
	}
	return nil
}

func (d walkingRecordDatabase) AllUpsert(ws walking_record.WalkingRecords) error {
	var es []WalkingRecordEntity
	for _, w := range ws {
		es = append(es, NewWalkingRecordEntity(w))
	}

	err := d.db.Save(&es).Error
	if err != nil {
		return model_errors.NewInfrastructureError(err.Error())
	}
	return nil
}

func (d walkingRecordDatabase) FindByFilter(
	filter walking_record.WalkingRecordFilter,
	order walking_record.WalkingRecordOrder,
) (walking_record.WalkingRecords, error) {
	filterE := NewWalkingRecordFilterEntity(filter)
	orderE := NewWalkingRecordOrderEntity(order)

	var es []WalkingRecordEntity
	err := d.db.
		Scopes(
			filterE.FilterQuery,
			orderE.OrderQuery,
		).Find(&es).Error
	if err != nil {
		return nil, model_errors.NewInfrastructureError(err.Error())
	}

	var ws walking_record.WalkingRecords
	for _, e := range es {
		ws = append(ws, e.ToModel())
	}

	return ws, nil
}
