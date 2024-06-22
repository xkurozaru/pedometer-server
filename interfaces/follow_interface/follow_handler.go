package follow_interface

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
	follow_application "github.com/xkurozaru/pedometer-server/application/follow"
	user_application "github.com/xkurozaru/pedometer-server/application/user"
	"github.com/xkurozaru/pedometer-server/domain/user"
	"github.com/xkurozaru/pedometer-server/interfaces"
)

type FollowHandler interface {
	PostFollow(w http.ResponseWriter, r *http.Request)
	DeleteFollow(w http.ResponseWriter, r *http.Request)
}

type followHandler struct {
	followApplicationService follow_application.FollowApplicationService
	userApplicationService   user_application.UserApplicationService
}

func NewFollowHandler(
	followApplicationService follow_application.FollowApplicationService,
	userApplicationService user_application.UserApplicationService,
) FollowHandler {
	return followHandler{
		followApplicationService: followApplicationService,
		userApplicationService:   userApplicationService,
	}
}

func (h followHandler) PostFollow(w http.ResponseWriter, r *http.Request) {
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

	var req PostFollowRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.followApplicationService.Follow(u.UserID(), user.UserID(req.FollowedUserID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h followHandler) DeleteFollow(w http.ResponseWriter, r *http.Request) {
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

	var req DeleteFollowRequest
	if err := schema.NewDecoder().Decode(&req, r.URL.Query()); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.followApplicationService.Unfollow(u.UserID(), user.UserID(req.FollowedUserID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
