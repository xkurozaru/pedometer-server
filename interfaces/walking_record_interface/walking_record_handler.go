package walking_record_interface

import (
	"encoding/json"
	"log/slog"
	"net/http"

	user_application "github.com/xkurozaru/pedometer-server/application/user"
	walking_record_application "github.com/xkurozaru/pedometer-server/application/walking_record"
	"github.com/xkurozaru/pedometer-server/domain/common"
	"github.com/xkurozaru/pedometer-server/interfaces"
)

type WalkingRecordHandler interface {
	GetWalkingRecord(w http.ResponseWriter, r *http.Request)
	PostWalkingRecord(w http.ResponseWriter, r *http.Request)
}

type walkingRecordHandler struct {
	userApplicationService          user_application.UserApplicationService
	walkingRecordApplicationService walking_record_application.WalkingRecordApplicationService
}

func NewWalkingRecordHandler(
	userApplicationService user_application.UserApplicationService,
	walkingRecordApplicationService walking_record_application.WalkingRecordApplicationService,
) WalkingRecordHandler {
	return walkingRecordHandler{
		userApplicationService:          userApplicationService,
		walkingRecordApplicationService: walkingRecordApplicationService,
	}
}

func (h walkingRecordHandler) GetWalkingRecord(w http.ResponseWriter, r *http.Request) {
	token, err := interfaces.GetToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	u, err := h.userApplicationService.FetchUserByToken(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	dto, err := h.walkingRecordApplicationService.FetchFriendsWeeklyWalkingRecordDistance(u.UserID(), common.DateTimeNow())
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := NewGetWalkingRecordResponse(dto)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h walkingRecordHandler) PostWalkingRecord(w http.ResponseWriter, r *http.Request) {
	token, err := interfaces.GetToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	u, err := h.userApplicationService.FetchUserByToken(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	var req PostWalkingRecordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	command := []walking_record_application.ApplyWalkingRecordCommand{}
	for _, record := range req.Records {
		command = append(command, walking_record_application.ApplyWalkingRecordCommand{
			Date:     record.Date,
			Distance: record.Distance,
		})
	}

	err = h.walkingRecordApplicationService.ApplyWalkingRecords(u.UserID(), command)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
