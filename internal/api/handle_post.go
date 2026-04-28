package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"net/url"
)

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
