package service

import (
	"cool-service/internal/models"
	"cool-service/internal/repository"
	"time"
)

type NoteService struct {
	repo repository.NoteRepository
}

func NewNoteService(r repository.NoteRepository) *NoteService {
	return &NoteService{repo: r}
}

func (s *NoteService) Create(note *models.Note) error {
	now := time.Now()
	note.CreatedAt = now
	note.UpdatedAt = &now
	return s.repo.Create(note)
}

func (s *NoteService) Update(note *models.Note) error {
	now := time.Now()
	note.UpdatedAt = &now
	return s.repo.Update(note)
}

func (s *NoteService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *NoteService) GetByID(id int) (*models.Note, error) {
	return s.repo.GetByID(id)
}

func (s *NoteService) GetAll() ([]models.Note, error) {
	return s.repo.GetAll()
}
