package main

import (
	"fmt"

	"github.com/IlhamRanggaKurniawan/Teamers.git/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	server := server.NewServer()

	err := server.ListenAndServe()

	if err != nil {
		panic(fmt.Sprintf("cannot start the app: %s", err))
	}
}