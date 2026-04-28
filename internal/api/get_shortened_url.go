package api

import (
	"errors"
	"fmt"
	"goshort/internal/store"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
)

func handleGet(store store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "code")
		data, err := store.Get(r.Context(), code)
		if err != nil {
			if errors.Is(err, redis.Nil) {
				slog.Info("code not found", "code", code)
				sendJSON(
					w,
					Response{Error: "code not found"},
					http.StatusNotFound,
				)
			}

			slog.Info("something went wrong", "code", code)
			sendJSON(
				w,
				Response{Error: "something went wrong"},
				http.StatusInternalServerError,
			)
			return
		}
		fmt.Printf("Redirecting to %s\n", data)
		http.Redirect(w, r, data, http.StatusFound)
	}
}
