package class

import (
	"encoding/json"
	"net/http"

	"github.com/IlhamRanggaKurniawan/Teamers.git/internal/database/entity"
	"github.com/IlhamRanggaKurniawan/Teamers.git/internal/utils"
)

type Handler struct {
	classService ClassService
}

type Input struct {
	Name     string        `json:"name"`
	Students []entity.User `json:"students"`
}

func NewHandler(classService ClassService) Handler {
	return Handler{classService: classService}
}

func (h *Handler) CreateClass(w http.ResponseWriter, r *http.Request) {
	var input Input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	user, err := utils.DecodeAccessToken(r)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnauthorized)
		return
	}

	class, err := h.classService.CreateClass(user.Id, input.Name)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	go func() {
		utils.GenerateAndSetAccessToken(w, user.Id, user.Username, user.Email, "Class Leader", class.Id)
		utils.GenerateAndSetRefreshToken(w, user.Id, user.Username, user.Email, "Class Leader", class.Id)
	}()

	utils.SuccessResponse(w, class)
}

func (h *Handler) FindClass(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetNumberPathParam(r, "id")

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	class, err := h.classService.FindClass(id)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	utils.SuccessResponse(w, class)
}

func (h *Handler) AddStudents(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetNumberPathParam(r, "id")

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

	class, err := h.classService.AddStudents(id, input.Students)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, class)
}
