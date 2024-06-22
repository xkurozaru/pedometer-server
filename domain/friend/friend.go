package friend

import (
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

func (f Friend) IsRequesting() bool {
	return f.status.isRequesting()
}
func (f Friend) IsRequested() bool {
	return f.status.isRequested()
}
func (f Friend) IsEstablished() bool {
	return f.status.isEstablished()
}

func (f Friend) Establish() Friends {
	return Friends{
		newFriend(f.userID, f.friendUserID, FriendStatusEstablished),
		newFriend(f.friendUserID, f.userID, FriendStatusEstablished),
	}
}

type Friends []Friend

func NewFriendPair(userID user.UserID, friendUserID user.UserID) Friends {
	return Friends{
		newFriend(userID, friendUserID, FriendStatusRequesting),
		newFriend(friendUserID, userID, FriendStatusRequested),
	}
}
