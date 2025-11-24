package main

import (
	"cool-service/internal/handlers"
	"cool-service/internal/service"
	"cool-service/internal/repository"
	"cool-service/internal/models"
	"log"
)

func main() {
	// init DB
	// init Repo
	// init Service
	// init Handler
	// run http
	db := repository.ConnectPostgres()
	db.AutoMigrate(&models.Note{})

	repo := repository.NewNoteRepository(db)
	service := service.NewNoteService(repo)
	handler := handlers.NewHandler(service)

	router := handler.InitRoutes()

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
