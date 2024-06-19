package record

import (
	"net/http"

	"github.com/xkurozaru/pedometer-server/dependency/registry"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	registry := registry.NewRegistry()

	switch r.Method {
	case http.MethodGet:
		registry.NewWalkingRecordHandler().GetWalkingRecord(w, r)
	case http.MethodPost:
		registry.NewWalkingRecordHandler().PostWalkingRecord(w, r)
	case http.MethodOptions:
		w.Write([]byte(""))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
