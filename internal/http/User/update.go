package user

import (
	"context"
	"net/http"
)

func (r Router) UpdatePassword(userId string, w http.ResponseWriter, rq *http.Request) error {
	password := rq.FormValue("password")

	return r.useCase.UpdatePassword(userId, password, context.Background())
}

func (r Router) UpdateLogin(userId string, w http.ResponseWriter, rq *http.Request) error {
	login := rq.FormValue("login")

	return r.useCase.UpdateLogin(userId, login, context.Background())
}

func (r Router) UpdateEmail(userId string, w http.ResponseWriter, rq *http.Request) error {
	email := rq.FormValue("email")

	return r.useCase.UpdateEmail(userId, email, context.Background())
}
