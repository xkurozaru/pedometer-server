package walking_record_application

import (
	"fmt"
	"log/slog"

	"github.com/xkurozaru/pedometer-server/domain/common"
	"github.com/xkurozaru/pedometer-server/domain/friend"
	"github.com/xkurozaru/pedometer-server/domain/user"
	"github.com/xkurozaru/pedometer-server/domain/walking_record"
)

type WalkingRecordApplicationService interface {
	ApplyWalkingRecords(userID user.UserID, command []ApplyWalkingRecordCommand) error
	FetchFriendsWeeklyWalkingRecordDistance(userID user.UserID, month common.DateTime) ([]WalkingRecordDistanceDTO, error)
}

type walkingRecordApplicationService struct {
	friendRepository        friend.FriendRepository
	userRepository          user.UserRepository
	walkingRecordRepository walking_record.WalkingRecordRepository
}

func NewWalkingRecordApplicationService(
	friendRepository friend.FriendRepository,
	userRepository user.UserRepository,
	walkingRecordRepository walking_record.WalkingRecordRepository,
) WalkingRecordApplicationService {
	return walkingRecordApplicationService{
		friendRepository:        friendRepository,
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
	slog.Info("NewWalkingRecord", "records", records)

	err := s.walkingRecordRepository.UpsertAll(records)
	if err != nil {
		return fmt.Errorf("UpsertAll: %w", err)
	}
	slog.Info("UpsertAll")

	return nil
}

func (s walkingRecordApplicationService) FetchFriendsWeeklyWalkingRecordDistance(userID user.UserID, date common.DateTime) ([]WalkingRecordDistanceDTO, error) {
	friends, err := s.friendRepository.FindFriendUsers(userID, friend.FriendStatusEstablished)
	if err != nil {
		return nil, fmt.Errorf("FindFriendUsers: %w", err)
	}
	slog.Info("FindFriendUsers", "friends", friends)

	if len(friends) == 0 {
		return []WalkingRecordDistanceDTO{}, nil
	}

	filter := walking_record.NewWalkingRecordFilter(friends.UserIDs(), date.StartOfWeek(), date.EndOfWeek())
	slog.Info("NewWalkingRecordFilter", "filter", filter)

	records, err := s.walkingRecordRepository.FindByFilter(filter, walking_record.WalkingRecordOrderUserIDAsc)
	if err != nil {
		return nil, fmt.Errorf("FindByFilter: %w", err)
	}
	slog.Info("FindByFilter", "records", records)

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
