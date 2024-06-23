package friend_interface

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gorilla/schema"
	friend_application "github.com/xkurozaru/pedometer-server/application/friend"
	user_application "github.com/xkurozaru/pedometer-server/application/user"
	"github.com/xkurozaru/pedometer-server/domain/user"
	"github.com/xkurozaru/pedometer-server/interfaces"
)

type FriendHandler interface {
	GetFriend(w http.ResponseWriter, r *http.Request)
	PostFriend(w http.ResponseWriter, r *http.Request)
	PatchFriend(w http.ResponseWriter, r *http.Request)
	DeleteFriend(w http.ResponseWriter, r *http.Request)
}

type friendHandler struct {
	userApplicationService   user_application.UserApplicationService
	friendApplicationService friend_application.FriendApplicationService
}

func NewFriendHandler(
	userApplicationService user_application.UserApplicationService,
	friendApplicationService friend_application.FriendApplicationService,
) FriendHandler {
	return friendHandler{
		userApplicationService:   userApplicationService,
		friendApplicationService: friendApplicationService,
	}
}

func (h friendHandler) GetFriend(w http.ResponseWriter, r *http.Request) {
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

	dto, err := h.friendApplicationService.FetchFriendList(u.UserID())
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := NewGetFriendResponse(dto)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h friendHandler) PostFriend(w http.ResponseWriter, r *http.Request) {
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

	var req PostFriendRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.friendApplicationService.RegisterFriendRequest(u.UserID(), user.UserID(req.FriendUserID))
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h friendHandler) PatchFriend(w http.ResponseWriter, r *http.Request) {
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

	var req PatchFriendRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.friendApplicationService.AcceptFriendRequest(u.UserID(), user.UserID(req.FriendUserID))
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h friendHandler) DeleteFriend(w http.ResponseWriter, r *http.Request) {
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

	var req DeleteFriendRequest
	if err := schema.NewDecoder().Decode(&req, r.URL.Query()); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.friendApplicationService.RemoveFriend(u.UserID(), user.UserID(req.FriendUserID))
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
