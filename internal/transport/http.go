package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type TodoRouter interface {
	GetOne(string, http.ResponseWriter, *http.Request)
	GetAll(string, http.ResponseWriter, *http.Request)
	Delete(string, http.ResponseWriter, *http.Request)
	Add(string, http.ResponseWriter, *http.Request)
	Update(string, http.ResponseWriter, *http.Request)
}

type UserRouter interface {
	Get(string, http.ResponseWriter, *http.Request)
	Login(http.ResponseWriter, *http.Request) (string, error)
	Registration(http.ResponseWriter, *http.Request)
	Delete(string, http.ResponseWriter, *http.Request)
	UpdateLogin(string, http.ResponseWriter, *http.Request)
	UpdatePassword(string, http.ResponseWriter, *http.Request)
	UpdateEmail(string, http.ResponseWriter, *http.Request)
}

type Session interface {
	Add(func(http.ResponseWriter, *http.Request) (string, error)) http.HandlerFunc
	Check(func(string, http.ResponseWriter, *http.Request)) http.HandlerFunc
}

type Logger interface {
	Log(next http.Handler) http.Handler
}

func Listen(tr TodoRouter, ur UserRouter, session Session, log Logger, port string) error {
	r := chi.NewRouter()

	r.Use(log.Log)
	r.Use(middleware.Recoverer)
	r.Use(middleware.CleanPath)
	r.Use(middleware.GetHead)

	r.Post("/v1/auth", session.Add(ur.Login))

	r.Put("/v1/registration", ur.Registration)

	r.Get("/v1/user", session.Check(ur.Get))
	r.Get("/v1/todo", session.Check(tr.GetAll))
	r.Get("/v1/todo/{id}", session.Check(tr.GetOne))

	r.Put("/v1/todo", session.Check(tr.Add))

	r.Patch("/v1/user/password", session.Check(ur.UpdatePassword))
	r.Patch("/v1/user/login", session.Check(ur.UpdateLogin))
	r.Patch("/v1/user/email", session.Check(ur.UpdateEmail))
	r.Patch("/v1/todo", session.Check(tr.Update))

	r.Delete("/v1/todo/{id}", session.Check(tr.Delete))
	r.Delete("/v1/user", session.Check(ur.Delete))

	return http.ListenAndServe(port, r)
}
