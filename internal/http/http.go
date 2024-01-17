package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type TodoRouter interface {
	Delete(http.ResponseWriter, *http.Request)
	Add(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
}

type UserRouter interface {
	Login(http.ResponseWriter, *http.Request)
	Registration(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
	UpdateLogin(http.ResponseWriter, *http.Request)
	UpdatePassword(http.ResponseWriter, *http.Request)
	UpdateEmail(http.ResponseWriter, *http.Request)
}

type TemplateRouter interface {
	GetUser(http.ResponseWriter, *http.Request)
	GetTodoOne(http.ResponseWriter, *http.Request)
	GetTodoAll(http.ResponseWriter, *http.Request)
	Main(http.ResponseWriter, *http.Request)
	Login(http.ResponseWriter, *http.Request)
	Registration(http.ResponseWriter, *http.Request)
}

func Listen(tr TodoRouter, ur UserRouter, tmR TemplateRouter, port string) error {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.CleanPath)
	r.Use(middleware.GetHead)

	r.Get("/", tmR.Main)
	r.Get("/auth", tmR.Login)
	r.Get("/registration", tmR.Registration)

	r.Post("/auth", ur.Login)

	r.Put("/registration", ur.Registration)

	r.Get("/user", tmR.GetUser)
	r.Get("/todo", tmR.GetTodoAll)
	r.Get("/todo/{id}", tmR.GetTodoOne)

	r.Put("/todo", tr.Add)

	r.Patch("/user/password", ur.UpdatePassword)
	r.Patch("/user/login", ur.UpdateLogin)
	r.Patch("/user/email", ur.UpdateEmail)
	r.Patch("/todo", tr.Update)

	r.Delete("/todo/{id}", tr.Delete)
	r.Delete("/user", ur.Delete)

	return http.ListenAndServe(port, r)
}
