package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHandler(db map[string]string) http.Handler {

	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)

	// Grupo de API
	r.Route("/api", func(r chi.Router) {
		r.Post("/shorten", handlePost(db))
		r.Get("/{code}", handleGet(db))
	})

	return r
}
