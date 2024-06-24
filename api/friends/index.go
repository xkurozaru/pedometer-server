package friends

import (
	"net/http"

	"github.com/xkurozaru/pedometer-server/dependency/registry"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	handler := registry.NewRegistry().NewFriendHandler()

	switch r.Method {
	case http.MethodGet:
		handler.GetFriend(w, r)
	case http.MethodPost:
		handler.PostFriend(w, r)
	case http.MethodPatch:
		handler.PatchFriend(w, r)
	case http.MethodDelete:
		handler.DeleteFriend(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
