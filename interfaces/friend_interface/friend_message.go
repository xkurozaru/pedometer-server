package friend_interface

import friend_application "github.com/xkurozaru/pedometer-server/application/friend"

type GetFriendRequest struct{}
type GetFriendResponse struct {
	Friends []struct {
		UserID   string `json:"user_id"`
		Username string `json:"username"`
	} `json:"friends"`
	Requested []struct {
		UserID   string `json:"user_id"`
		Username string `json:"username"`
	} `json:"requested"`
	Requesting []struct {
		UserID   string `json:"user_id"`
		Username string `json:"username"`
	} `json:"requesting"`
}

func NewGetFriendResponse(dto friend_application.FriendListDTO) GetFriendResponse {
	res := GetFriendResponse{}
	for _, f := range dto.Friends {
		res.Friends = append(res.Friends, struct {
			UserID   string `json:"user_id"`
			Username string `json:"username"`
		}{
			UserID:   f.FriendUserID,
			Username: f.FriendUsername,
		})
	}
	for _, f := range dto.Requested {
		res.Requested = append(res.Requested, struct {
			UserID   string `json:"user_id"`
			Username string `json:"username"`
		}{
			UserID:   f.FriendUserID,
			Username: f.FriendUsername,
		})
	}
	for _, f := range dto.Requesting {
		res.Requesting = append(res.Requesting, struct {
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
