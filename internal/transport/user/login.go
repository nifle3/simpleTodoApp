package user

import (
	"context"
	"fmt"
	"net/http"
)

// @Summary	Log in
// @Tags		Auth service
// @Accept		json
// @Produce	json
// @Param		email		formData	string	true	"user's email"
// @Param		password	formData	string	true	"user's password"
// @Failure	400			{object}	error
// @Failure	401			{object}	error
// @Failure	500			{object}	error
// @Success	200
// @Router		/v1/auth [post]
func (r Router) Login(w http.ResponseWriter, rq *http.Request) (string, error) {
	if !(rq.Form.Has("email") && rq.Form.Has("password")) {
		err := fmt.Errorf("form has not enough parametr")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return "", err
	}
	email := rq.PostFormValue("email")
	password := rq.PostFormValue("password")

	userID, err := r.useCase.Login(email, password, context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return "", err
	}

	w.WriteHeader(http.StatusOK)
	return userID, nil
}
