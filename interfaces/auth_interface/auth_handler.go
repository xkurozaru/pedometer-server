package auth_interface

import (
	"encoding/json"
	"net/http"

	"github.com/HackU-team04/Pedometer/application"
)

type AuthHandler interface {
	PostAuth(w http.ResponseWriter, r *http.Request)
}

type authHandler struct {
	authApplicationService application.AuthApplicationService
}

func NewAuthHandler(
	authApplicationService application.AuthApplicationService,
) AuthHandler {
	return authHandler{
		authApplicationService: authApplicationService,
	}
}

func (h authHandler) PostAuth(w http.ResponseWriter, r *http.Request) {
	var req PostAuthRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := h.authApplicationService.Login(req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := PostAuthResponse{
		Token: token,
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}
