package walking_record

import (
	"github.com/xkurozaru/pedometer-server/domain/common"
	"github.com/xkurozaru/pedometer-server/domain/user"
)

type WalkingRecordFilter struct {
	UserIDs   []user.UserID
	StartDate common.DateTime
	EndDate   common.DateTime
}

func NewWalkingRecordFilter(
	userIDs []user.UserID,
	startDate common.DateTime,
	endDate common.DateTime,
) WalkingRecordFilter {
	return WalkingRecordFilter{
		UserIDs:   userIDs,
		StartDate: startDate,
		EndDate:   endDate,
	}
}
