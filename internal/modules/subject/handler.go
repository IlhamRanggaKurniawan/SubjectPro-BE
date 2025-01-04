package subject

import (
	"encoding/json"
	"net/http"

	"github.com/IlhamRanggaKurniawan/Teamers.git/internal/utils"
)

type Handler struct {
	subjectService SubjectService
}

type Input struct {
	Name string `json:"name"`
}

func NewHandler(subjectService SubjectService) Handler {
	return Handler{subjectService: subjectService}
}

func (h *Handler) CreateSubject(w http.ResponseWriter, r *http.Request) {
	classId, err := utils.GetNumberPathParam(r, "classId")

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

	subject, err := h.subjectService.CreateSubject(input.Name, classId)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, subject)
}

func (h *Handler) FindAllSubjects(w http.ResponseWriter, r *http.Request) {
	classId, err := utils.GetNumberPathParam(r, "classId")

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	subjects, err := h.subjectService.FindAllSubjects(classId)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, subjects)
}

func (h *Handler) DeleteSubject(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetNumberPathParam(r, "id")

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	err = h.subjectService.DeleteSubject(id)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusNotFound)
		return 
	}

	response := map[string]string{
		"Message": "Delete subject success",
	}

	utils.SuccessResponse(w, response)
}