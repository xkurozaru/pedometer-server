package record

import (
	"net/http"

	"github.com/xkurozaru/pedometer-server/dependency/registry"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	handler := registry.NewRegistry().NewWalkingRecordHandler()

	switch r.Method {
	case http.MethodGet:
		handler.GetWalkingRecord(w, r)
	case http.MethodPost:
		handler.PostWalkingRecord(w, r)
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
