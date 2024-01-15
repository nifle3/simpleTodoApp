package user

import (
	"context"
	"net/http"
	"todoApp/internal/domain"
)

func (r Router) Login(w http.ResponseWriter, rq *http.Request) (domain.User, error) {
	email := rq.PostFormValue("email")
	password := rq.PostFormValue("password")

	ctx := context.Background()

	return r.useCase.Login(email, password, ctx)
}
