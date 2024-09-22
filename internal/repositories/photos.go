package repositories

import (
	"database/sql"
	"test-backend/internal/models"
	"time"
)

type Photo interface {
	GetPhotos() ([]models.Photo, error)
	CreatePhoto(photo *models.Photo) error
	UpdatePhoto(photo *models.Photo) error
	DeletePhoto(id uint) error
}

type PhotoRepository struct {
	DB *sql.DB
}

func NewPhotoRepository(DB *sql.DB) *PhotoRepository {
	return &PhotoRepository{DB: DB}
}

func (r *PhotoRepository) GetPhotos() ([]models.Photo, error) {
	rows, err := r.DB.Query("SELECT * FROM photos")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var photos []models.Photo
	for rows.Next() {
		var photo models.Photo
		if err := rows.Scan(&photo.Id, &photo.Title, &photo.Description, &photo.Url, &photo.CreatedAt); err != nil {
			return nil, err
		}
		photos = append(photos, photo)
	}
	return photos, nil
}

func (r *PhotoRepository) CreatePhoto(photo *models.Photo) error {
	photo.CreatedAt = time.Now()

	_, err := r.DB.Exec("INSERT INTO photos (title, description, url, created_at) VALUES ($1, $2, $3, $4)", photo.Title, photo.Description, photo.Url, photo.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *PhotoRepository) UpdatePhoto(photo *models.Photo) error {
	photo.UpdatedAt = time.Now()

	_, err := r.DB.Exec("UPDATE photos SET title = ?, description = ?,  url = ? WHERE id = ?", photo.Title, photo.Description, photo.Url, photo.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r *PhotoRepository) DeletePhoto(id uint) error {
	_, err := r.DB.Exec("DELETE FROM photos WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
