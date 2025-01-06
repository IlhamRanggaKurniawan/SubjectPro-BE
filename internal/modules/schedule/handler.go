package schedule

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/IlhamRanggaKurniawan/Teamers.git/internal/utils"
)

type Handler struct {
	scheduleService ScheduleService
}

type Input struct {
	Day       string    `json:"day"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}

func NewHandler(scheduleService ScheduleService) Handler {
	return Handler{scheduleService: scheduleService}
}

func (h *Handler) CreateSchedule(w http.ResponseWriter, r *http.Request) {
	subjectId, err := utils.GetNumberPathParam(r, "subjectId")

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	var input Input

	err = json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	schedule, err := h.scheduleService.CreateSchedule(input.Day, subjectId, input.StartTime, input.EndTime)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, schedule)
}

func (h *Handler) FindAllScheduleByDay(w http.ResponseWriter, r *http.Request) {
	var input Input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	subjectId, err  := utils.GetNumberPathParam(r, "subjectId")

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	schedules, err := h.scheduleService.FindAllByDay(input.Day, subjectId)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	utils.SuccessResponse(w, schedules)
}

func(h *Handler) DeleteSchedule(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetNumberPathParam(r, "id")

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	err = h.scheduleService.DeleteSchedule(id)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return 
	}

	response := map[string]string{
		"message": "Schedule deleted",
	}

	utils.SuccessResponse(w, response)
}
