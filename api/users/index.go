package users

import (
	"net/http"

	"github.com/xkurozaru/pedometer-server/dependency/registry"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	registry := registry.NewRegistry()

	switch r.Method {
	case http.MethodGet:
		registry.NewUserHandler().GetUser(w, r)
	case http.MethodPost:
		registry.NewUserHandler().PostUser(w, r)
	case http.MethodOptions:
		w.Write([]byte(""))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
