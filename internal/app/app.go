package app

import (
	"context"
	"log/slog"
	"os"
	"todoApp/internal/config"
	"todoApp/internal/services/todo"
	"todoApp/internal/services/user"
	"todoApp/internal/storage/mongo"
	http "todoApp/internal/transport"
	jwttoken "todoApp/internal/transport/jwtToken"
	tr "todoApp/internal/transport/todo"
	ur "todoApp/internal/transport/user"
	"todoApp/pkg/middleware"
)

func Start() {
	cfg := config.MustLoad()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	logger.Info("App start", slog.Any("config is", cfg))

	db, err := mongo.New(cfg.Db.Port, context.Background())
	if err != nil {
		logger.LogAttrs(context.Background(), slog.LevelError, "DB:", slog.String("err", err.Error()))
		os.Exit(1)
	}

	userUseCase := user.NewUseCase(db)
	todoUseCase := todo.NewUseCase(db)

	userRouter := ur.New(userUseCase)
	todoRouter := tr.New(todoUseCase)

	jwt := jwttoken.New(cfg.JwtKey)
	middlewareLogger := middleware.New(logger)

	if err := http.Listen(todoRouter, userRouter, jwt, middlewareLogger, cfg.HttpServer.Port); err != nil {
		logger.LogAttrs(context.Background(), slog.LevelError, "HTTP:", slog.String("err", err.Error()))
	}
}
