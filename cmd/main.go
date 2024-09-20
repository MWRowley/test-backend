package main

import (
	"fmt"
	"net/http"

	"test-backend/internal/db"
	"test-backend/internal/handlers"
	"test-backend/internal/repositories"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	userRepo := repositories.NewUserRepository(db.DB)
	handlers.Handler(r, userRepo)

	db.Init()

	fmt.Println("Starting server on :8000")

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Error(err)
	}
}
