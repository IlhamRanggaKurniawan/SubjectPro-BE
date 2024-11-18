package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Server struct {
	port int
}

func NewServer() *http.Server{
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	newServer := &Server{
		port: port,
	}

	server := &http.Server{
		Addr: fmt.Sprintf(":%d", port),
		Handler: newServer.RegisterRoutes(),
	}

	return server
}