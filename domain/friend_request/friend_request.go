package friend_request

import (
	"github.com/xkurozaru/pedometer-server/domain/common"
	model_errors "github.com/xkurozaru/pedometer-server/domain/errors"
	"github.com/xkurozaru/pedometer-server/domain/user"
)

type FriendRequestID common.NanoID

type FriendRequest struct {
	friendRequestID FriendRequestID
	fromUserID      user.UserID
	toUserID        user.UserID
	status          FriendRequestStatus
}

func NewFriendRequest(fromUserID user.UserID, toUserID user.UserID) FriendRequest {
	return newFriendRequest(FriendRequestID(common.NewNanoID()), fromUserID, toUserID, FriendRequestStatusPending)
}
func RecreateFriendRequest(friendRequestID FriendRequestID, fromUserID user.UserID, toUserID user.UserID, status FriendRequestStatus) FriendRequest {
	return newFriendRequest(friendRequestID, fromUserID, toUserID, status)
}
func newFriendRequest(friendRequestID FriendRequestID, fromUserID user.UserID, toUserID user.UserID, status FriendRequestStatus) FriendRequest {
	return FriendRequest{
		friendRequestID: friendRequestID,
		fromUserID:      fromUserID,
		toUserID:        toUserID,
		status:          status,
	}
}

func (r FriendRequest) FriendRequestID() FriendRequestID {
	return r.friendRequestID
}
func (r FriendRequest) FromUserID() user.UserID {
	return r.fromUserID
}
func (r FriendRequest) ToUserID() user.UserID {
	return r.toUserID
}
func (r FriendRequest) Status() FriendRequestStatus {
	return r.status
}

func (r FriendRequest) Accept() (FriendRequest, error) {
	status, err := r.status.toAccept()
	if err != nil {
		return r, model_errors.NewInvalidError(err.Error())
	}
	r.status = status
	return r, nil
}
func (r FriendRequest) Reject() (FriendRequest, error) {
	status, err := r.status.toReject()
	if err != nil {
		return r, model_errors.NewInvalidError(err.Error())
	}
	r.status = status
	return r, nil
}
