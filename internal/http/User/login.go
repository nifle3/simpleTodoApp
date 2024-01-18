package user

import (
	"context"
	"net/http"
)

func (r Router) Login(w http.ResponseWriter, rq *http.Request) {
	email := rq.PostFormValue("email")
	password := rq.PostFormValue("password")

	if _, err := r.useCase.Login(email, password, context.Background()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
