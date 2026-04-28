package api

import (
	"encoding/json"
	"goshort/internal/store"
	"log/slog"
	"net/http"
	"net/url"
)

func handlePost(store store.Store) http.HandlerFunc {
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

		code, err := store.SaveShortenedURL(r.Context(), body.URL)

		if err != nil {
			slog.Error("failed to save shortened url", "error", err)
			sendJSON(
				w,
				Response{Error: "something went wrong"},
				http.StatusInternalServerError,
			)
			return
		}

		sendJSON(
			w,
			Response{Data: code},
			http.StatusCreated,
		)
	}
}
