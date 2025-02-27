package user_interface

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/xkurozaru/pedometer-server/domain/user"

	user_application "github.com/xkurozaru/pedometer-server/application/user"
	"github.com/xkurozaru/pedometer-server/interfaces"
)

type UserHandler interface {
	GetUser(w http.ResponseWriter, r *http.Request)
	PostUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	userApplicationService user_application.UserApplicationService
}

func NewUserHandler(
	userApplicationService user_application.UserApplicationService,
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

	var req GetUserRequest
	if err := schema.NewDecoder().Decode(&req, r.URL.Query()); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var u user.User
	switch req.UserID {
	case "":
		u, err = h.userApplicationService.FetchUserByToken(token)
	default:
		u, err = h.userApplicationService.FetchUserByUserID(user.UserID(req.UserID))
	}
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := GetUserResponse{
		UserID:   string(u.UserID()),
		Username: u.Username(),
	}

	w.Header().Set("Content-Type", "application/json")
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

	err := h.userApplicationService.RegisterUser(req.Email, req.Password, user.UserID(req.UserID), req.Username)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h userHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	token, err := interfaces.GetToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	u, err := h.userApplicationService.FetchUserByToken(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	err = h.userApplicationService.Delete(u)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
