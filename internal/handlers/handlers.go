package handlers

import (
	"test-backend/internal/repositories"

	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux, userRepo *repositories.UserRepository, postRepo *repositories.PostRepository, photoRepo *repositories.PhotoRepository) {
	r.Use(chimiddle.StripSlashes)

	userHandler := &UserHandler{Repo: userRepo}

	r.Route("/api", func(router chi.Router) {
		router.Get("/users", userHandler.GetUsers)
		router.Post("/users", userHandler.CreateUser)
		router.Put("/users/{id}", userHandler.UpdateUser)
		router.Delete("/users/{id}", userHandler.DeleteUser)
	})

	postHandler := &PostHandler{Repo: postRepo}

	r.Route("/api", func(router chi.Router) {
		router.Get("/posts", postHandler.GetPosts)
		router.Post("/posts", postHandler.CreatePost)
		router.Put("/posts/{id}", postHandler.UpdatePost)
		router.Delete("/posts/{id}", postHandler.DeletePost)
	})

	photoHandler := &PhotoHandler{Repo: photoRepo}

	r.Route("/api", func(router chi.Router) {
		router.Get("/photos", photoHandler.GetPhotos)
		router.Post("/photos", photoHandler.CreatePhoto)
		router.Put("/photos/{id}", photoHandler.UpdatePhoto)
		router.Delete("/photos/{id}", photoHandler.DeletePhoto)
	})
}
