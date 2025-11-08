package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "MORSE: ", log.Ldate|log.Ltime)
	server := server.New(logger)
	logger.Println("Server starting on :8080")
	err := server.Run()
	if err != nil {
		logger.Fatal("Server failed to start: ", err)
	}
}
