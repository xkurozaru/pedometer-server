package friend_interface

import friend_application "github.com/xkurozaru/pedometer-server/application/friend"

type GetFriendRequest struct {
	Status string `schema:"status"`
}
type GetFriendResponse struct {
	Friends []struct {
		UserID   string `json:"user_id"`
		Username string `json:"username"`
	} `json:"friends"`
}

func NewGetFriendResponse(dto []friend_application.FriendDTO) GetFriendResponse {
	res := GetFriendResponse{}
	for _, f := range dto {
		res.Friends = append(res.Friends, struct {
			UserID   string `json:"user_id"`
			Username string `json:"username"`
		}{
			UserID:   f.FriendUserID,
			Username: f.FriendUsername,
		})
	}
	return res
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
