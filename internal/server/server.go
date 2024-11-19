package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/IlhamRanggaKurniawan/Teamers.git/internal/database"
	"gorm.io/gorm"
)

type Server struct {
	port int
	DB *gorm.DB
}

func NewServer() *http.Server{
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	DB := database.NewDB()

	newServer := &Server{
		port: port,
		DB: DB,
	}

	server := &http.Server{
		Addr: fmt.Sprintf(":%d", port),
		Handler: newServer.RegisterRoutes(),
	}

	return server
}