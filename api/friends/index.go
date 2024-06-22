package friends

import (
	"net/http"

	"github.com/xkurozaru/pedometer-server/dependency/registry"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	registry := registry.NewRegistry()

	switch r.Method {
	case http.MethodGet:
		registry.NewFriendHandler().GetFriend(w, r)
	case http.MethodPost:
		registry.NewFriendHandler().PostFriend(w, r)
	case http.MethodPatch:
		registry.NewFriendHandler().PatchFriend(w, r)
	case http.MethodDelete:
		registry.NewFriendHandler().DeleteFriend(w, r)
	case http.MethodOptions:
		w.Write([]byte(""))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
