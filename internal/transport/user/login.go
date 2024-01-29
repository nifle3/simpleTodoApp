package user

import (
	"context"
	"net/http"
)

func (r Router) Login(w http.ResponseWriter, rq *http.Request) (string, error) {
	email := rq.PostFormValue("email")
	password := rq.PostFormValue("password")

	user, err := r.useCase.Login(email, password, context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return "", err
	}

	w.WriteHeader(http.StatusOK)
	return user.ID, nil
}
