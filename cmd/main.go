package main

import (
	"fmt"
	"net/http"

	"test-backend/configs"
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
	photoRepo := repositories.NewPhotoRepository(db.DB)

	seed.SeedUsers(userRepo)
	seed.SeedPosts(postRepo)
	seed.SeedPhotos(photoRepo)

	handlers.Handler(r, userRepo, postRepo, photoRepo)

	configs.LoadServerConfig()

	port := configs.ServerConfig.Port
	fmt.Println("Starting server on port: ", port)

	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
