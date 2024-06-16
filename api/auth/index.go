package auth

import (
	"net/http"

	"github.com/HackU-team04/Pedometer/dependency/registry"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	registry := registry.NewRegistry()

	switch r.Method {
	case http.MethodPost:
		registry.NewAuthHandler().PostAuth(w, r)
	case http.MethodOptions:
		w.Write([]byte(""))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
