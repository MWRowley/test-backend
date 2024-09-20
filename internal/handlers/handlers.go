package handlers

import (
	"test-backend/internal/repositories"

	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux, userRepo *repositories.UserRepository) {
	r.Use(chimiddle.StripSlashes)

	userHandler := &UserHandler{Repo: userRepo}

	r.Route("/api", func(router chi.Router) {
		router.Get("/users", userHandler.GetUsers)
		router.Post("/users", userHandler.CreateUser)
		router.Put("/users/{id}", userHandler.UpdateUser)
		router.Delete("/users/{id}", userHandler.DeleteUser)
	})
}
