package walking_record

import "github.com/xkurozaru/pedometer-server/domain/common"

type WalkingRecordFilter struct {
	UserIDs   []string
	StartDate common.DateTime
	EndDate   common.DateTime
}

func NewWalkingRecordFilter(
	userIDs []string,
	startDate common.DateTime,
	endDate common.DateTime,
) WalkingRecordFilter {
	return WalkingRecordFilter{
		UserIDs:   userIDs,
		StartDate: startDate,
		EndDate:   endDate,
	}
}
