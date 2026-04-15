package main

import (
	"log/slog"
	"os"
)

func main() {
	if err := run(); err != nil {
		slog.Error("Failed to Execute code", "Error", err)
		os.Exit(1)
	}
	slog.Info("All system offline")
}

func run() error {

	return nil

}
