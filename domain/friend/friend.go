package friend

import (
	"github.com/xkurozaru/pedometer-server/domain/friend_request"
	"github.com/xkurozaru/pedometer-server/domain/user"
)

type Friend struct {
	userID       user.UserID
	friendUserID user.UserID
}

func NewFriend(userID user.UserID, friendUserID user.UserID) Friend {
	return Friend{
		userID:       userID,
		friendUserID: friendUserID,
	}
}

func (f Friend) UserID() user.UserID {
	return f.userID
}
func (f Friend) FriendUserID() user.UserID {
	return f.friendUserID
}

type Friends []Friend

func NewFriendPairs(req friend_request.FriendRequest) Friends {
	return Friends{
		NewFriend(req.FromUserID(), req.ToUserID()),
		NewFriend(req.ToUserID(), req.FromUserID()),
	}
}
