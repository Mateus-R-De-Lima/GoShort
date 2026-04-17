package api

//go get github.com/go-chi/chi/v5 -> Comando para instalar pacote chi
import (
	"encoding/json"
	"log/slog"
	"math/rand"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type PostBody struct {
	URL string `json:"url"`
}

type Response struct {
	Error string `json:"error,omitempty"`
	Data  string `json:"data,omitempty"`
}

func NewHandler(db map[string]string) http.Handler {

	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)

	// Rotas
	r.Post("/shorten", handlePost(db))
	r.Get("/{code}", handleGet(db))

	return r
}
func sendJSON(w http.ResponseWriter, resp Response, status int) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error("failed to marshal json data", "error", err)
		sendJSON(
			w,
			Response{Error: "something went wrong"},
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		slog.Error("failed to write json data", "error", err)
		return
	}
}

const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func genCode() string {
	const n = 8
	byts := make([]byte, n)
	for i := range n {
		byts[i] = characters[rand.Intn(len(characters))]
	}
	return string(byts)
}

func handlePost(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body PostBody
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			slog.Error("failed to decode request body", "error", err)
			sendJSON(
				w,
				Response{Error: "invalid request body"},
				http.StatusBadRequest,
			)
			return
		}
		if _, err := url.Parse(body.URL); err != nil {
			slog.Error("invalid url", "error", err)
			sendJSON(
				w,
				Response{Error: "invalid url"},
				http.StatusBadRequest,
			)
			return
		}

		code := genCode()
		db[code] = body.URL
		sendJSON(
			w,
			Response{Data: code},
			http.StatusCreated,
		)
	}
}

func handleGet(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "code")
		data, ok := db[code]
		if !ok {
			slog.Info("code not found", "code", code)
			sendJSON(
				w,
				Response{Error: "code not found"},
				http.StatusNotFound,
			)
			return
		}
		http.Redirect(w, r, data, http.StatusFound)
	}
}
