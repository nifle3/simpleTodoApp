package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type TodoRouter interface {
}

type UserRouter interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type SessionsRouter interface {
	CheckSession(next http.Handler) http.Handler
	AddSession(next http.Handler) http.Handler
}

type ErrorRouter interface {
	CheckError(next http.Handler) http.Handler
}

func Listen(tr TodoRouter, ur UserRouter, sr SessionsRouter, er ErrorRouter, port string) {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(er.CheckError)
	r.Use(middleware.CleanPath)
	r.Use(middleware.GetHead)

	r.Get("/", nil)
	r.Get("/auth", nil)
	r.Get("/registration", nil)

	r.Post("/auth", nil)

	r.Put("/registration", nil)

	r.Group(func(r chi.Router) {
		r.Use(sr.CheckSession)

		r.Get("/user", nil)
		r.Get("/todo", nil)
		r.Get("/todo/{id}", nil)

		r.Put("/todo", nil)

		r.Patch("/user/password", nil)
		r.Patch("/user/login", nil)
		r.Patch("/user/email", nil)
		r.Patch("/todo", nil)

		r.Delete("/todo/{id}", nil)
		r.Delete("/user", nil)
	})

	http.ListenAndServe(port, r)
}
