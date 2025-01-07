package server

import (
	"net/http"

	"github.com/IlhamRanggaKurniawan/Teamers.git/internal/middleware"
	"github.com/IlhamRanggaKurniawan/Teamers.git/internal/modules/class"
	"github.com/IlhamRanggaKurniawan/Teamers.git/internal/modules/schedule"
	"github.com/IlhamRanggaKurniawan/Teamers.git/internal/modules/subject"
	"github.com/IlhamRanggaKurniawan/Teamers.git/internal/modules/task"
	"github.com/IlhamRanggaKurniawan/Teamers.git/internal/modules/user"
)

func roleMiddleware(handler func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(handler)
}

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	middlewares := middleware.CreateStack(middleware.CORSMiddleware, middleware.AuthMiddelware)

	userRepository := user.NewRepo(s.DB)
	userService := user.NewService(userRepository)
	userHandler := user.NewHandler(userService)

	classRepository := class.NewRepo(s.DB)
	classService := class.NewService(classRepository, userRepository)
	classHandler := class.NewHandler(classService)

	subjectRepository := subject.NewRepo(s.DB)
	subjectService := subject.NewService(subjectRepository)
	subjectHandler := subject.NewHandler(subjectService)

	scheduleRepository := schedule.NewRepo(s.DB)
	scheduleService := schedule.NewService(scheduleRepository)
	scheduleHandler := schedule.NewHandler(scheduleService)

	taskRepository := task.NewRepo(s.DB)
	taskService := task.NewService(taskRepository)
	taskHandler := task.NewHandler(taskService)

	mux.HandleFunc("POST /v1/auth", userHandler.Register)
	mux.HandleFunc("POST /v1/auth/login", userHandler.Login)
	mux.HandleFunc("POST /v1/auth/logout", userHandler.Logout)
	mux.HandleFunc("GET /v1/auth/token", userHandler.GetToken)

	mux.HandleFunc("POST /v1/class", classHandler.CreateClass)
	mux.HandleFunc("GET /v1/class/{id}", classHandler.FindClass)
	// mux.HandleFunc("GET /v1/class/{id}/{day}", classHandler.FindClassWithSchedule)
	mux.Handle("PATCH /v1/class/{id}", roleMiddleware(classHandler.AddStudents))

	mux.Handle("POST /v1/subject/{classId}", roleMiddleware(subjectHandler.CreateSubject))
	mux.HandleFunc("GET /v1/subject/{classId}", subjectHandler.FindAllSubjects)
	mux.HandleFunc("GET /v1/subject/{classId}/{day}/{deadline}", subjectHandler.FindAllSubjectByDeadline)
	mux.HandleFunc("GET /v1/subject/{classId}/{day}", subjectHandler.FindAllSubjectByDay)
	mux.Handle("DELETE /v1/subject/{id}", roleMiddleware(subjectHandler.DeleteSubject))

	mux.Handle("POST /v1/schedule/{subjectId}", roleMiddleware(scheduleHandler.CreateSchedule))
	mux.HandleFunc("GET /v1/schedule/{subjectId}/{day}", scheduleHandler.FindAllScheduleByDay)
	mux.Handle("DELETE /v1/schedule/{id}", roleMiddleware(scheduleHandler.DeleteSchedule))

	mux.Handle("POST /v1/task/{subjectId}", roleMiddleware(taskHandler.CreateTask))
	mux.HandleFunc("GET /v1/task/{subjectId}/{deadline}", taskHandler.FindAllTaskByDeadline)
	mux.Handle("DELETE /v1/task/{id}", roleMiddleware(taskHandler.DeleteTask))
	return middlewares(mux)
}
