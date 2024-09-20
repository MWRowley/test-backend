package main

import (
	"fmt"
	"net/http"

	"test-backend/internal/db"
	"test-backend/internal/handlers"
	"test-backend/internal/repositories"
	"test-backend/internal/seed"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)

	db.Init()

	var r *chi.Mux = chi.NewRouter()

	userRepo := repositories.NewUserRepository(db.DB)
	postRepo := repositories.NewPostRepository(db.DB)

	seed.SeedUsers(userRepo)
	seed.SeedPosts(postRepo)

	handlers.Handler(r, userRepo)

	fmt.Println("Starting server on :8000")

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
