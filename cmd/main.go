package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/ridwankustanto/family-tree-tracker/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	log.Println("Listening on port 8080...")
	server.Run()
}
