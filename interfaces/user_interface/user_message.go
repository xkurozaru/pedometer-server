package user_interface

type GetUserRequest struct{}
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
