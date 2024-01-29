package user

import (
	"context"
	"net/http"
)

func (r Router) UpdatePassword(userID string, w http.ResponseWriter, rq *http.Request) {
	password := rq.FormValue("password")

	if err := r.useCase.UpdatePassword(userID, password, context.Background()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (r Router) UpdateLogin(userID string, w http.ResponseWriter, rq *http.Request) {
	login := rq.FormValue("login")

	if err := r.useCase.UpdateLogin(userID, login, context.Background()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (r Router) UpdateEmail(userID string, w http.ResponseWriter, rq *http.Request) {
	email := rq.FormValue("email")

	if err := r.useCase.UpdateEmail(userID, email, context.Background()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
