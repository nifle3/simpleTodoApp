package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type TodoRouter interface {
	Delete(http.ResponseWriter, *http.Request) error
	Add(http.ResponseWriter, *http.Request) error
	Update(http.ResponseWriter, *http.Request) error
	GetAll(http.ResponseWriter, *http.Request) error
	GetOne(http.ResponseWriter, *http.Request) error
}

type UserRouter interface {
}

type SessionsRouter interface {
	CheckSession(next http.Handler) http.Handler
	AddSession(next http.Handler) http.Handler
}

type ErrorRouter interface {
	CheckError(next ErrorHandler) http.HandlerFunc
}

type ErrorHandler func(http.ResponseWriter, *http.Request) error

func Listen(tr TodoRouter, ur UserRouter, sr SessionsRouter, er ErrorRouter, port string) {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
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
		r.Get("/todo", er.CheckError(tr.GetAll))
		r.Get("/todo/{id}", er.CheckError(tr.GetOne))

		r.Put("/todo", er.CheckError(tr.Add))

		r.Patch("/user/password", nil)
		r.Patch("/user/login", nil)
		r.Patch("/user/email", nil)
		r.Patch("/todo", er.CheckError(tr.Update))

		r.Delete("/todo/{id}", er.CheckError(tr.Delete))
		r.Delete("/user", nil)
	})

	http.ListenAndServe(port, r)
}
