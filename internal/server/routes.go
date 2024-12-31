package server

import (
	"net/http"

	"github.com/IlhamRanggaKurniawan/Teamers.git/internal/middleware"
	"github.com/IlhamRanggaKurniawan/Teamers.git/internal/modules/user"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	middlewares := middleware.CreateStack(middleware.AuthMiddelware, middleware.CORSMiddleware)

	userRepository := user.NewRepo(s.DB)
	userService := user.NewService(userRepository)
	userHandler := user.NewHandler(userService)

	mux.HandleFunc("POST /v1/auth", userHandler.Register)
	mux.HandleFunc("POST /v1/auth/login", userHandler.Login)
	mux.HandleFunc("POST /v1/auth/logout", userHandler.Logout)
	mux.HandleFunc("GET /v1/auth/token", userHandler.GetToken)
	
	return middlewares(mux)
}
