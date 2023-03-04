package http

import (
	"github.com/endyApina/exercise-admin-computer/db"
	"github.com/endyApina/exercise-admin-computer/server/http/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
)

func MountServer(store db.DataStore) *chi.Mux {
	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)

	///implement handlers
	httpHandler := handlers.NewHttpHandler(store)

	//health check
	router.Get("/health", httpHandler.TestHealth)

	return router
}
