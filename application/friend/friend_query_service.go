package friend_application

import "github.com/xkurozaru/pedometer-server/domain/user"

type FriendQueryService interface {
	GetFriendList(userID user.UserID) (FriendListDTO, error)
}
