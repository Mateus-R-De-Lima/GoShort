package api

//go get github.com/go-chi/chi/v5 -> Comando para instalar pacote chi
import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type PostBody struct {
	URL string `json:"url"`
}

type Response struct {
	Error string `json:"error,omitempty"`
	URL   string `json:"url,omitempty"`
}

func NewHandler() http.Handler {

	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)

	// Rotas
	r.Post("/shorten", handlePost)
	r.Get("/{id}", handleGet)

	return r
}

func handlePost(w http.ResponseWriter, r *http.Request) {

}

func handleGet(w http.ResponseWriter, r *http.Request) {

}
