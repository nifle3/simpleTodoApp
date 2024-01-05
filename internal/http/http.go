package http

import (
	"net/http"

	"github.com/go-chi/chi"
)

func Listen() {
	r := chi.NewRouter()

	r.Get("/", nil)
	r.Post("/login", nil)

	r.Group(func(r chi.Router) {
		r.Use()

		r.Get("/todo", nil)
		r.Get("/todo/{id}", nil)
		r.Put("/todo", nil)
		r.Patch("/route/{id}", nil)
	})

	http.ListenAndServe(":8080", r)
}
