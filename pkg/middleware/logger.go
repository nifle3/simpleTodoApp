package middleware

import (
	"context"
	"log/slog"
	"net/http"
)

type Logger struct {
	log *slog.Logger
}

func New(log *slog.Logger) Logger {
	return Logger{
		log: log,
	}
}

func (l Logger) Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l.log.LogAttrs(context.Background(), slog.LevelInfo, "Request", slog.String("uri", r.RequestURI))
		l.log.LogAttrs(context.Background(), slog.LevelInfo, "Request", slog.String("ip", r.RemoteAddr))
		l.log.LogAttrs(context.Background(), slog.LevelInfo, "Request", slog.Any("Header", r.Header))
		l.log.LogAttrs(context.Background(), slog.LevelInfo, "Request", slog.Any("Body", r.Body))

		next.ServeHTTP(w, r)
	})
}
