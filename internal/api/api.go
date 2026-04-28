package api

import (
	"goshort/internal/store"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHandler(store store.Store) http.Handler {

	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)

	// Grupo de API
	r.Route("/api", func(r chi.Router) {
		r.Post("/shorten", handlePost(store))
		r.Get("/{code}", handleGet(store))
	})

	return r
}
