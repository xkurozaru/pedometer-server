package friend_interface

type GetFriendRequest struct {
	Status string `schema:"status"`
}
type GetFriendResponse struct {
	Friends []struct {
		UserID   string `json:"user_id"`
		Username string `json:"username"`
	} `json:"friends"`
}

type PostFriendRequest struct {
	FriendUserID string `json:"friend_user_id"`
}
type PostFriendResponse struct{}

type PatchFriendRequest struct {
	FriendUserID string `json:"friend_user_id"`
}
type PatchFriendResponse struct{}

type DeleteFriendRequest struct {
	FriendUserID string `schema:"friend-user-id,required"`
}
type DeleteFriendResponse struct{}
