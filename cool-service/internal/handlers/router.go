package handlers

import (
	"cool-service/internal/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type Handler struct {
	Service *service.NoteService
}

func NewHandler(s *service.NoteService) *Handler {
	return &Handler{Service: s}
}

func (h *Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()
	api1 := r.PathPrefix("/api/v1").Subrouter()

	notes := api1.PathPrefix("/notes").Subrouter()
	notes.Path("/").HandlerFunc(h.GetNotes).Methods(http.MethodGet)
	notes.Path("/by-id").HandlerFunc(h.GetNoteById).
		Methods(http.MethodGet)
	notes.Path("/delete").HandlerFunc(h.DeleteNote).
		Methods(http.MethodDelete)
	notes.Path("/update").HandlerFunc(h.UpdateNote).
		Methods(http.MethodPut)
	notes.Path("/create").HandlerFunc(h.CreateNote).
		Methods(http.MethodPost)
		
	return r
}