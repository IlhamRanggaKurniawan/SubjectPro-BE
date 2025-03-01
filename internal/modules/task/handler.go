package task

import (
	"encoding/json"
	"net/http"

	"github.com/IlhamRanggaKurniawan/Teamers.git/internal/utils"
)

type Handler struct {
	taskService TaskService
}

type Input struct {
	TaskType string    `json:"taskType"`
	Note     string    `json:"note"`
	Deadline string `json:"deadline"`
}

func NewHandler(taskService TaskService) Handler {
	return Handler{taskService: taskService}
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
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

	task, err := h.taskService.CreateTask(subjectId, input.TaskType, input.Note, input.Deadline)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, task)
}

func (h *Handler) FindAllTaskByDeadline(w http.ResponseWriter, r *http.Request) {
	subjectId, err := utils.GetNumberPathParam(r, "subjectId")

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	deadline, err := utils.GetStringPathParam(r, "deadline")

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	tasks, err := h.taskService.FindAllByDeadline(deadline, subjectId)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusNotFound)
		return
	}

	utils.SuccessResponse(w, tasks)
}

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetNumberPathParam(r, "id")

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	err = h.taskService.DeleteTask(id)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	response := map[string]string{
		"message": "Task is deleted",
	}

	utils.SuccessResponse(w, response)
}
