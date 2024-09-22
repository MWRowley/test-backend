package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"test-backend/internal/models"
	"test-backend/internal/repositories"

	"github.com/gorilla/mux"
)

type PhotoHandler struct {
	Repo *repositories.PhotoRepository
}

func (h *PhotoHandler) GetPhotos(w http.ResponseWriter, r *http.Request) {
	photos, err := h.Repo.GetPhotos()
	if err != nil {
		log.Println("Error getting photos", err)
		http.Error(w, "Error getting photos", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(photos)
}

func (h *PhotoHandler) CreatePhoto(w http.ResponseWriter, r *http.Request) {
	var photo models.Photo
	if err := json.NewDecoder(r.Body).Decode(&photo); err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}
	err := h.Repo.CreatePhoto(&photo)
	if err != nil {
		log.Println("Error creating photo", err)
		http.Error(w, "Error creating photo", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(photo)
}

func (h *PhotoHandler) UpdatePhoto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	photoId, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid Photo ID", http.StatusBadRequest)
		return
	}
	var photo models.Photo
	if err := json.NewDecoder(r.Body).Decode(&photo); err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	photo.Id = photoId
	err = h.Repo.UpdatePhoto(&photo)
	if err != nil {
		log.Println("Error updating photo", err)
		http.Error(w, "Error updating photo", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(photo)
}

func (h *PhotoHandler) DeletePhoto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	photoId, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid Photo ID", http.StatusBadRequest)
		return
	}

	err = h.Repo.DeletePhoto(uint(photoId))
	if err != nil {
		log.Println("Error deleting photo", err)
		http.Error(w, "Error deleting photo", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)
}
