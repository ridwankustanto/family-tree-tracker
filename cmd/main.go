package main

import (
	"log"

	"github.com/kelseyhightower/envconfig"
	"github.com/ridwankustanto/family-tree-tracker/server"
)

type Config struct {
	DatabaseURL string `envconfig:"DATABASE_URL"`
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Listening on port 8080...")
	server.Run()
}
