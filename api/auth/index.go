package auth

import (
	"net/http"

	"github.com/xkurozaru/pedometer-server/dependency/registry"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	handler := registry.NewRegistry().NewAuthHandler()

	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("<h1>登録が完了しました</h1>"))
	case http.MethodPost:
		handler.PostAuth(w, r)
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
