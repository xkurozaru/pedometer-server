package follow_interface

type PostFollowRequest struct {
	FollowedUserID string `json:"followed_user_id"`
}
type PostFollowResponse struct{}

type DeleteFollowRequest struct {
	FollowedUserID string `schema:"followed-user-id,required"`
}
type DeleteFollowResponse struct{}
