package users

import (
	"net/http"

	"github.com/xkurozaru/pedometer-server/dependency/registry"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	handler := registry.NewRegistry().NewUserHandler()

	switch r.Method {
	case http.MethodGet:
		handler.GetUser(w, r)
	case http.MethodPost:
		handler.PostUser(w, r)
	case http.MethodDelete:
		handler.DeleteUser(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
