package friend

import (
	model_errors "github.com/xkurozaru/pedometer-server/domain/errors"
	"github.com/xkurozaru/pedometer-server/domain/user"
)

type Friend struct {
	userID       user.UserID
	friendUserID user.UserID
	status       FriendStatus
}

func RecreateFriend(userID user.UserID, friendUserID user.UserID, status FriendStatus) Friend {
	return newFriend(userID, friendUserID, status)
}
func newFriend(userID user.UserID, friendUserID user.UserID, status FriendStatus) Friend {
	return Friend{
		userID:       userID,
		friendUserID: friendUserID,
		status:       status,
	}
}

func (f Friend) UserID() user.UserID {
	return f.userID
}
func (f Friend) FriendUserID() user.UserID {
	return f.friendUserID
}
func (f Friend) Status() FriendStatus {
	return f.status
}

func (f Friend) CanEstablish() bool {
	return f.status.isRequested()
}

func (f Friend) Establish() (Friends, error) {
	if !f.CanEstablish() {
		return Friends{f}, model_errors.NewAlreadyExistsError(f.friendUserID)
	}
	return Friends{
		newFriend(f.userID, f.friendUserID, FriendStatusEstablished),
		newFriend(f.friendUserID, f.userID, FriendStatusEstablished),
	}, nil
}

type Friends []Friend

func NewFriendPair(userID user.UserID, friendUserID user.UserID) Friends {
	return Friends{
		newFriend(userID, friendUserID, FriendStatusRequesting),
		newFriend(friendUserID, userID, FriendStatusRequested),
	}
}
