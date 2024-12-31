package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/IlhamRanggaKurniawan/Teamers.git/internal/database/entity"
	"github.com/IlhamRanggaKurniawan/Teamers.git/internal/utils"
)

type Handler struct {
	userService UserService
}

type Input struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	ConfPassword string `json:"confPassword"`
}

type AuthenticationRes struct {
	User entity.User `json:"user"`
	AccessToken string `json:"accessToken"`
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

	fmt.Println(user)

	utils.SuccessResponse(w, user)

	// accessToken, err := utils.GenerateAccessToken(user.Id)
}