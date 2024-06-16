package auth

import (
	"net/http"

	"github.com/xkurozaru/pedometer-server/dependency/registry"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	registry := registry.NewRegistry()

	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("<h1>登録が完了しました</h1>"))
	case http.MethodPost:
		registry.NewAuthHandler().PostAuth(w, r)
	case http.MethodOptions:
		w.Write([]byte(""))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
