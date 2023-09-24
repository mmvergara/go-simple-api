package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/mmvergara/go-simple-api/handlers"
)

func loadPostRoutes(router chi.Router) {
	postHandler := &handlers.Post{}

	router.Post("/", postHandler.Create)
	router.Get("/", postHandler.List)
	router.Get("/{id}", postHandler.GetByID)
	router.Put("/{id}", postHandler.UpdateByID)
	router.Delete("/{id}", postHandler.DeleteByID)
}
