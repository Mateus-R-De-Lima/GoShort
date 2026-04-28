package api

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

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
		fmt.Printf("Redirecting to %s\n", data)
		http.Redirect(w, r, data, http.StatusFound)
	}
}
