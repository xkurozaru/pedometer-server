package user_interface

import (
	"encoding/json"
	"net/http"

	"github.com/xkurozaru/pedometer-server/application"
	"github.com/xkurozaru/pedometer-server/interfaces"
)

type UserHandler interface {
	GetUser(w http.ResponseWriter, r *http.Request)
	PostUser(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	userApplicationService application.UserApplicationService
}

func NewUserHandler(
	userApplicationService application.UserApplicationService,
) UserHandler {
	return userHandler{
		userApplicationService: userApplicationService,
	}
}

func (h userHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	token, err := interfaces.GetToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u, err := h.userApplicationService.GetUser(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := GetUserResponse{
		UserID:   u.UserID(),
		Username: u.Username(),
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h userHandler) PostUser(w http.ResponseWriter, r *http.Request) {
	var req PostUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.userApplicationService.RegisterUser(req.Email, req.Password, req.UserID, req.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
