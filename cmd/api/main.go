package main

import (
	"goshort/internal/api"
	"goshort/internal/store"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	if err := run(); err != nil {
		slog.Error("Failed to Execute code", "Error", err)
		os.Exit(1)
	}
	slog.Info("All system offline")
}

func run() error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer rdb.Close()

	store := store.NewStore(rdb)

	handler := api.NewHandler(store)

	s := http.Server{
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
		Handler:      handler,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil

}
