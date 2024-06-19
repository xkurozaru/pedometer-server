package walking_record_application

import (
	"fmt"

	"github.com/xkurozaru/pedometer-server/domain/common"
	"github.com/xkurozaru/pedometer-server/domain/user"
	"github.com/xkurozaru/pedometer-server/domain/walking_record"
)

type WalkingRecordApplicationService interface {
	ApplyWalkingRecords(userID string, command []ApplyWalkingRecordCommand) error
	FetchFriendsWeeklyWalkingRecordDistance(userID string, month common.DateTime) ([]WalkingRecordDistanceDTO, error)
}

type walkingRecordApplicationService struct {
	userRepository          user.UserRepository
	walkingRecordRepository walking_record.WalkingRecordRepository
}

func NewWalkingRecordApplicationService(
	userRepository user.UserRepository,
	walkingRecordRepository walking_record.WalkingRecordRepository,
) WalkingRecordApplicationService {
	return walkingRecordApplicationService{
		userRepository:          userRepository,
		walkingRecordRepository: walkingRecordRepository,
	}
}

func (s walkingRecordApplicationService) ApplyWalkingRecords(userID string, command []ApplyWalkingRecordCommand) error {
	records := walking_record.WalkingRecords{}
	for _, c := range command {
		date, err := common.DateTimeFromString(c.Date, common.HyphenDateFormat)
		if err != nil {
			return fmt.Errorf("DateTimeFromString: %w", err)
		}
		record := walking_record.NewWalkingRecord(userID, date, c.Distance)
		records = append(records, record)
	}

	err := s.walkingRecordRepository.AllUpsert(records)
	if err != nil {
		return fmt.Errorf("AllUpsert: %w", err)
	}

	return nil
}

// NOTE: フレンドの距離を取得する予定だが、全ユーザーの距離を取得するように仮置き
func (s walkingRecordApplicationService) FetchFriendsWeeklyWalkingRecordDistance(userID string, date common.DateTime) ([]WalkingRecordDistanceDTO, error) {
	// TODO: フレンドに置き換える
	users, err := s.userRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("FindAll: %w", err)
	}

	filter := walking_record.NewWalkingRecordFilter([]string{}, date.StartOfWeek(), date.EndOfWeek())

	records, err := s.walkingRecordRepository.FindByFilter(filter, walking_record.WalkingRecordOrderUserIDAsc)
	if err != nil {
		return nil, fmt.Errorf("FindByFilter: %w", err)
	}

	dto := []WalkingRecordDistanceDTO{}
	total := records.TotalUserDistance()
	for _, u := range users {
		dto = append(dto, WalkingRecordDistanceDTO{
			UserID:   u.UserID(),
			Username: u.Username(),
			Distance: total[u.UserID()],
		})
	}

	return dto, nil
}
