package walking_record_interface

import walking_record_application "github.com/xkurozaru/pedometer-server/application/walking_record"

type PostWalkingRecordRequest struct {
	Records []struct {
		Date     string `json:"date"`
		Distance int    `json:"distance"`
	} `json:"records"`
}
type PostWalkingRecordResponse struct{}

type GetWalkingRecordRequest struct{}
type GetWalkingRecordResponse struct {
	FriendDistances []struct {
		UserID   string `json:"user_id"`
		Username string `json:"username"`
		Distance int    `json:"distance"`
	} `json:"friend_distances"`
}

func NewGetWalkingRecordResponse(dto []walking_record_application.WalkingRecordDistanceDTO) GetWalkingRecordResponse {
	res := GetWalkingRecordResponse{}
	for _, d := range dto {
		res.FriendDistances = append(res.FriendDistances, struct {
			UserID   string `json:"user_id"`
			Username string `json:"username"`
			Distance int    `json:"distance"`
		}{
			UserID:   d.UserID,
			Username: d.Username,
			Distance: d.Distance,
		})
	}
	return res
}
