package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/IlhamRanggaKurniawan/Teamers.git/internal/utils"
)

type Handler struct {
	userService UserService
}

type Input struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	ConfPassword string `json:"confPassword"`
}

func NewHandler(userService UserService) Handler {
	return Handler{
		userService: userService,
	}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var input Input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	if input.Password != input.ConfPassword {
		utils.ErrorResponse(w, fmt.Errorf("password doesn't match"), http.StatusBadRequest)
		return
	}

	user, err := h.userService.Register(input.Username, input.Email, input.Password)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	_, err = utils.GenerateAndSetRefreshToken(w, user.Id, user.Username, user.Email, user.Role, user.ClassId)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	go func(){
		utils.GenerateAndSetAccessToken(w, user.Id, user.Username, user.Email, user.Role, user.ClassId)
	}()

	utils.SuccessResponse(w, user)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var input Input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	user, err := h.userService.Login(input.Email, input.Password)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	_, err = utils.GenerateAndSetRefreshToken(w, user.Id, user.Username, user.Email, user.Role, user.ClassId)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	go func(){
		utils.GenerateAndSetAccessToken(w, user.Id, user.Username, user.Email, user.Role, user.ClassId)
	}()

	utils.SuccessResponse(w, user)
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "AccessToken",
		Value:    "",
		Expires:  time.Now().Add(-1),
		Secure:   os.Getenv("APP_ENV") == "production",
		HttpOnly: true,
		Path:     "/",
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "RefreshToken",
		Value:    "",
		Expires:  time.Now().Add(-1),
		Secure:   os.Getenv("APP_ENV") == "production",
		HttpOnly: true,
		Path:     "/",
	})

	response := map[string]string{
		"Message": "Logout success",
	}

	utils.SuccessResponse(w, response)
}

func (h *Handler) GetToken(w http.ResponseWriter, r *http.Request) {
	user, err := utils.DecodeRefreshToken(r)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnauthorized)
		return
	}

	accessToken, err := utils.GenerateAndSetAccessToken(w, user.Id, user.Username, user.Email, user.Role, user.ClassId)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"accessToken": accessToken,
	}

	utils.SuccessResponse(w, response)
}
