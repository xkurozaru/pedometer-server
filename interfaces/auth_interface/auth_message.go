package auth_interface

type PostAuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type PostAuthResponse struct {
	Token string `json:"token"`
}
