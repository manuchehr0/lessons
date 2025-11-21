package main

import (
	"cool-service/internal/handlers"
	"log"
)

func main() {
	// init DB
	// init Repo
	// init Service
	// init Handler
	// run http

	h := handlers.NewHandler(":8080")

	log.Println("listening")
	h.InitRoutes()

}
