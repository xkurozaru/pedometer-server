package follow

import (
	"net/http"

	"github.com/xkurozaru/pedometer-server/dependency/registry"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	registry := registry.NewRegistry()

	switch r.Method {
	case http.MethodPost:
		registry.NewFollowHandler().PostFollow(w, r)
	case http.MethodDelete:
		registry.NewFollowHandler().DeleteFollow(w, r)
	case http.MethodOptions:
		w.Write([]byte(""))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
