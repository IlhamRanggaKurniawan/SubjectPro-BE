package server

import (
	"net/http"

	"github.com/IlhamRanggaKurniawan/Teamers.git/internal/middleware"
	"github.com/IlhamRanggaKurniawan/Teamers.git/internal/modules/class"
	"github.com/IlhamRanggaKurniawan/Teamers.git/internal/modules/subject"
	"github.com/IlhamRanggaKurniawan/Teamers.git/internal/modules/user"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	middlewares := middleware.CreateStack(middleware.CORSMiddleware, middleware.AuthMiddelware)

	userRepository := user.NewRepo(s.DB)
	userService := user.NewService(userRepository)
	userHandler := user.NewHandler(userService)

	classRepository := class.NewRepo(s.DB)
	classService := class.NewService(classRepository)
	classHandler := class.NewHandler(classService)

	subjectRepository := subject.NewRepo(s.DB)
	subjectService := subject.NewService(subjectRepository)
	subjectHandler := subject.NewHandler(subjectService)

	mux.HandleFunc("POST /v1/auth", userHandler.Register)
	mux.HandleFunc("POST /v1/auth/login", userHandler.Login)
	mux.HandleFunc("POST /v1/auth/logout", userHandler.Logout)
	mux.HandleFunc("GET /v1/auth/token", userHandler.GetToken)

	mux.HandleFunc("POST /v1/class", classHandler.CreateClass)
	mux.HandleFunc("GET /v1/class/{id}", classHandler.FindClass)
	mux.HandleFunc("PATCH /v1/class/{id}", classHandler.AddStudents)

	mux.HandleFunc("POST /v1/subject/{classId}", subjectHandler.CreateSubject)
	mux.HandleFunc("GET /v1/subject/{classId}", subjectHandler.FindAllSubjects)
	mux.HandleFunc("DELETE /v1/subject/{id}", subjectHandler.DeleteSubject)

	return middlewares(mux)
}
