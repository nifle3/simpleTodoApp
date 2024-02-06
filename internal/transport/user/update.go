package user

import (
	"context"
	"net/http"
)

//	@Summary	Update password
//	@Tags		User
//	@Accept		json
//	@Produce	json
//	@Param		password	formData	string	true	"new password"
//	@Failure	401			{object}	error
//	@Failure	500			{object}	error
//	@Success	200
//	@Router		/v1/user/password [patch]
func (r Router) UpdatePassword(userID string, w http.ResponseWriter, rq *http.Request) {
	password := rq.FormValue("password")

	if err := r.useCase.UpdatePassword(userID, password, context.Background()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

//	@Summary	Update login
//	@Tags		User
//	@Accept		json
//	@Produce	json
//
//	@Security	ApiKeyAuth
//
//	@Param		login	formData	string	true	"new login"
//	@Failure	401		{object}	error
//	@Failure	500		{object}	error
//	@Success	200
//	@Router		/v1/user/login [patch]
func (r Router) UpdateLogin(userID string, w http.ResponseWriter, rq *http.Request) {
	login := rq.FormValue("login")

	if err := r.useCase.UpdateLogin(userID, login, context.Background()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

//	@Summary	Update emaail
//	@Tags		User
//	@Accept		json
//	@Produce	json
//	@Security	ApiKeyAuth
//	@Param		email	formData	string	true	"new email"
//	@Failure	401		{object}	error
//	@Failure	500		{object}	error
//	@Success	200
//	@Router		/v1/user/email [patch]
func (r Router) UpdateEmail(userID string, w http.ResponseWriter, rq *http.Request) {
	email := rq.FormValue("email")

	if err := r.useCase.UpdateEmail(userID, email, context.Background()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
