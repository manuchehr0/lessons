package main

import (
	"cool-service/internal/handlers"
	"cool-service/internal/service"
	"cool-service/internal/repository"
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
	handler := handlers.NewHandler(":8080", service)
	handler.InitRoutes()

}
