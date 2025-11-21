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
	Port string
}

func NewHandler(port string) *Handler {
	return &Handler{Port: port}
}

func (h *Handler) InitRoutes() {
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

	if err := http.ListenAndServe(h.Port, r); err != nil {
		log.Fatal(err)
	}
}

func (h *Handler) GetNotes(w http.ResponseWriter, r *http.Request) {
	notes := map[string]string{
		"note": "some text",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(notes)
}
func (h *Handler) CreateNote(w http.ResponseWriter, r *http.Request) {
	note := models.Note{}

	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {

	}
	log.Println("new note:", note)

	note.Id = 123
	now := time.Now()
	note.CreatedAt = time.Now()
	note.UpdatedAt = &now
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(note)
}
func (h *Handler) UpdateNote(w http.ResponseWriter, r *http.Request) {
	notes := map[string]string{
		"note": "some text",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(notes)
}
func (h *Handler) DeleteNote(w http.ResponseWriter, r *http.Request) {
	notes := map[string]string{
		"note": "some text",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(notes)
}
func (h *Handler) GetNoteById(w http.ResponseWriter, r *http.Request) {
	notes := map[string]string{
		"note": "some text",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(notes)
}
