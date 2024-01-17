package app

import (
	"log/slog"
	"os"
	"todoApp/internal/config"
)

func Start() {
	cfg := config.MustLoad()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	logger.Info("App start", slog.Any("config is", cfg))

	// db = storage.New()

	// server = server.New()

	// server.Listen()
}
