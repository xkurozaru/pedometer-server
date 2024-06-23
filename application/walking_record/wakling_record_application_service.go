package walking_record_application

import (
	"fmt"

	"github.com/xkurozaru/pedometer-server/domain/common"
	"github.com/xkurozaru/pedometer-server/domain/user"
	"github.com/xkurozaru/pedometer-server/domain/walking_record"
)

type WalkingRecordApplicationService interface {
	ApplyWalkingRecords(userID user.UserID, command []ApplyWalkingRecordCommand) error
	FetchFriendsWeeklyWalkingRecordDistance(userID user.UserID, month common.DateTime) ([]WalkingRecordDistanceDTO, error)
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

func (s walkingRecordApplicationService) ApplyWalkingRecords(userID user.UserID, command []ApplyWalkingRecordCommand) error {
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
func (s walkingRecordApplicationService) FetchFriendsWeeklyWalkingRecordDistance(userID user.UserID, date common.DateTime) ([]WalkingRecordDistanceDTO, error) {
	// TODO: フレンドに置き換える
	friends, err := s.userRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("FindAll: %w", err)
	}

	friendIDs := []user.UserID{}
	for _, f := range friends {
		// NOTE: 自分自身は除外
		if f.UserID() == userID {
			continue
		}
		friendIDs = append(friendIDs, f.UserID())
	}

	filter := walking_record.NewWalkingRecordFilter(friendIDs, date.StartOfWeek(), date.EndOfWeek())

	records, err := s.walkingRecordRepository.FindByFilter(filter, walking_record.WalkingRecordOrderUserIDAsc)
	if err != nil {
		return nil, fmt.Errorf("FindByFilter: %w", err)
	}

	dto := []WalkingRecordDistanceDTO{}
	totalMap := records.TotalUserDistance()
	for _, f := range friends {
		dto = append(dto, WalkingRecordDistanceDTO{
			UserID:   string(f.UserID()),
			Username: f.Username(),
			Distance: totalMap[f.UserID()],
		})
	}

	return dto, nil
}
