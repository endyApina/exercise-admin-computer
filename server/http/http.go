package http

import (
	"github.com/endyApina/exercise-admin-computer/lib/idgenerator"

	"github.com/endyApina/exercise-admin-computer/db"
	"github.com/endyApina/exercise-admin-computer/server/http/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
)

func MountServer(store db.DataStore, idGenerator idgenerator.IdGenerator) *chi.Mux {
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
	httpHandler := handlers.NewHttpHandler(store, idGenerator)

	//health check
	router.Get("/health", httpHandler.TestHealth)
	router.Post("/computer/", httpHandler.CreateComputer)
	router.Get("/computer/", httpHandler.GetAllComputers)
	router.Get("/computer/{computer_id}", httpHandler.GetComputerByComputerID)
	router.Get("/computer/employee/{employee_abbreviation}", httpHandler.GetComputersByEmployeeName)
	router.Delete("/computer/{computer_id}", httpHandler.DeleteComputerByComputerID)
	router.Patch("/computer/", httpHandler.UpdateComputerAllocation)

	return router
}
