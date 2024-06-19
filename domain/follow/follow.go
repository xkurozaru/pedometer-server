package follow

type Follow struct {
	userID         string
	followedUserID string
}

func NewFollow(userID string, followedUserID string) Follow {
	return newFollow(userID, followedUserID)
}
func RecreateFollow(userID string, followedUserID string) Follow {
	return newFollow(userID, followedUserID)
}
func newFollow(userID string, followedUserID string) Follow {
	return Follow{
		userID:         userID,
		followedUserID: followedUserID,
	}
}

func (f Follow) UserID() string {
	return f.userID
}
func (f Follow) FollowedUserID() string {
	return f.followedUserID
}

type Follows []Follow

func (fs Follows) FollowedUserIDs() []string {
	ids := []string{}
	for _, f := range fs {
		ids = append(ids, f.FollowedUserID())
	}
	return ids
}
