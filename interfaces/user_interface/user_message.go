package user_interface

type GetUserRequest struct {
	UserID string `schema:"user-id"`
}
type GetUserResponse struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
}

type PostUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	UserID   string `json:"user_id"`
	Username string `json:"username"`
}
type PostUserResponse struct{}

type DeleteUserRequest struct{}
type DeleteUserResponse struct{}
