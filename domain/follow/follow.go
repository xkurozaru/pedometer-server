package follow

import "github.com/xkurozaru/pedometer-server/domain/user"

type Follow struct {
	userID         user.UserID
	followedUserID user.UserID
}

func NewFollow(userID user.UserID, followedUserID user.UserID) Follow {
	return newFollow(userID, followedUserID)
}
func RecreateFollow(userID user.UserID, followedUserID user.UserID) Follow {
	return newFollow(userID, followedUserID)
}
func newFollow(userID user.UserID, followedUserID user.UserID) Follow {
	return Follow{
		userID:         userID,
		followedUserID: followedUserID,
	}
}

func (f Follow) UserID() user.UserID {
	return f.userID
}
func (f Follow) FollowedUserID() user.UserID {
	return f.followedUserID
}

type Follows []Follow

func (fs Follows) FollowedUserIDs() []user.UserID {
	var ids []user.UserID
	for _, f := range fs {
		ids = append(ids, f.FollowedUserID())
	}
	return ids
}
