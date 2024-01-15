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
}

type UserRouter interface {
	Login(http.ResponseWriter, *http.Request) error
	Registration(http.ResponseWriter, *http.Request) error
	Delete(http.ResponseWriter, *http.Request) error
	UpdateLogin(http.ResponseWriter, *http.Request) error
	UpdatePassword(http.ResponseWriter, *http.Request) error
	UpdateEmail(http.ResponseWriter, *http.Request) error
}

type TemplateRouter interface {
	GetUser(http.ResponseWriter, *http.Request) error
	GetTodoOne(http.ResponseWriter, *http.Request) error
	GetTodoAll(http.ResponseWriter, *http.Request) error
	Main(http.ResponseWriter, *http.Request) error
	Login(http.ResponseWriter, *http.Request) error
	Registration(http.ResponseWriter, *http.Request) error
}

type SessionsRouter interface {
	CheckSession(next http.Handler) http.Handler
	AddSession(next http.Handler) http.Handler
}

type ErrorRouter interface {
	CheckError(next ErrorHandler) http.HandlerFunc
}

type ErrorHandler func(http.ResponseWriter, *http.Request) error

func Listen(tr TodoRouter, ur UserRouter, sr SessionsRouter, er ErrorRouter, tmR TemplateRouter, port string) {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.CleanPath)
	r.Use(middleware.GetHead)

	r.Get("/", er.CheckError(tmR.Main))
	r.Get("/auth", er.CheckError(tmR.Login))
	r.Get("/registration", er.CheckError(tmR.Registration))

	r.Post("/auth", er.CheckError(ur.Login))

	r.Put("/registration", er.CheckError(ur.Registration))

	r.Group(func(r chi.Router) {
		r.Use(sr.CheckSession)

		r.Get("/user", er.CheckError(tmR.GetUser))
		r.Get("/todo", er.CheckError(tmR.GetTodoAll))
		r.Get("/todo/{id}", er.CheckError(tmR.GetTodoOne))

		r.Put("/todo", er.CheckError(tr.Add))

		r.Patch("/user/password", er.CheckError(ur.UpdatePassword))
		r.Patch("/user/login", er.CheckError(ur.UpdateLogin))
		r.Patch("/user/email", er.CheckError(ur.UpdateEmail))
		r.Patch("/todo", er.CheckError(tr.Update))

		r.Delete("/todo/{id}", er.CheckError(tr.Delete))
		r.Delete("/user", er.CheckError(ur.Delete))
	})

	http.ListenAndServe(port, r)
}
