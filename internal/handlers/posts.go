package handlers

import (
	"test-backend/internal/models"
	"test-backend/internal/repositories"

	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type PostHandler struct {
	Repo *repositories.PostRepository
}

func (h *PostHandler) GetPosts(w http.ResponseWriter, r *http.Request) {
	post, err := h.Repo.GetPosts()
	if err != nil {
		log.Println("Error getting posts", err)
		http.Error(w, "Error getting posts", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

func (h *PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
	}

	post.Id = postId
	err = h.Repo.UpdatePost(&post)
	if err != nil {
		log.Println("Error updating post", err)
		http.Error(w, "Error updating post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func (h *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	err = h.Repo.DeletePost(uint(postId))
	if err != nil {
		log.Println("Error deleting post ", err)
		http.Error(w, "Error deleting post", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)
}
