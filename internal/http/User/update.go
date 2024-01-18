package user

import (
	"context"
	"net/http"
)

func (r Router) UpdatePassword(w http.ResponseWriter, rq *http.Request) {
	password := rq.FormValue("password")

	if err := r.useCase.UpdatePassword("", password, context.Background()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (r Router) UpdateLogin(w http.ResponseWriter, rq *http.Request) {
	login := rq.FormValue("login")

	if err := r.useCase.UpdateLogin("", login, context.Background()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (r Router) UpdateEmail(w http.ResponseWriter, rq *http.Request) {
	email := rq.FormValue("email")

	if err := r.useCase.UpdateEmail("", email, context.Background()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
