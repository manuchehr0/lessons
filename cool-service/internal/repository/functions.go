package repository

import (
	"cool-service/internal/models"
	"gorm.io/gorm"
)

type NoteRepository interface {
	Create(note *models.Note) error
	Update(note *models.Note) error
	Delete(id int) error
	GetByID(id int) (*models.Note, error)
	GetAll() ([]models.Note, error)
}

type noteRepository struct {
	db *gorm.DB
}

func NewNoteRepository(db *gorm.DB) NoteRepository {
	return &noteRepository{db: db}
}

func (r *noteRepository) Create(note *models.Note) error {
	return r.db.Create(note).Error
}

func (r *noteRepository) Update(note *models.Note) error {
	return r.db.Save(note).Error
}

func (r *noteRepository) Delete(id int) error {
	return r.db.Delete(&models.Note{}, id).Error
}

func (r *noteRepository) GetByID(id int) (*models.Note, error) {
	note := models.Note{}
	err := r.db.First(&note, id).Error
	return &note, err
}

func (r *noteRepository) GetAll() ([]models.Note, error) {
	notes := []models.Note{}
	err := r.db.Find(&notes).Error
	return notes, err
}

