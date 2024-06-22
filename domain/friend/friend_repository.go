package friend

import "github.com/xkurozaru/pedometer-server/domain/user"

type FriendRepository interface {
	Find(userID, friendUserID user.UserID) (Friend, error)
	Exists(userID, friendUserID user.UserID) (bool, error)
	UpsertAll(friends Friends) error
	Delete(userID, friendUserID user.UserID) error
}
