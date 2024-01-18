package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type TodoRouter interface {
	GetOne(http.ResponseWriter, *http.Request)
	GetAll(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
	Add(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
}

type UserRouter interface {
	Get(http.ResponseWriter, *http.Request)
	Login(http.ResponseWriter, *http.Request)
	Registration(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
	UpdateLogin(http.ResponseWriter, *http.Request)
	UpdatePassword(http.ResponseWriter, *http.Request)
	UpdateEmail(http.ResponseWriter, *http.Request)
}

func Listen(tr TodoRouter, ur UserRouter, port string) error {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.CleanPath)
	r.Use(middleware.GetHead)

	r.Post("/auth", ur.Login)

	r.Put("/registration", ur.Registration)

	r.Get("/user", ur.Get)
	r.Get("/todo", tr.GetAll)
	r.Get("/todo/{id}", tr.GetOne)

	r.Put("/todo", tr.Add)

	r.Patch("/user/password", ur.UpdatePassword)
	r.Patch("/user/login", ur.UpdateLogin)
	r.Patch("/user/email", ur.UpdateEmail)
	r.Patch("/todo", tr.Update)

	r.Delete("/todo/{id}", tr.Delete)
	r.Delete("/user", ur.Delete)

	return http.ListenAndServe(port, r)
}
