package user

import (
	"context"
	"todoApp/internal/domain"
)

type UseCase interface {
	Login(email, password string, ctx context.Context) (domain.User, error)
	Add(user domain.User, ctx context.Context) error
	Delete(userID string, ctx context.Context) error
	UpdatePassword(userID, password string, ctx context.Context) error
	UpdateEmail(userID, email string, ctx context.Context) error
	UpdateLogin(userID, login string, ctx context.Context) error
	Get(userID string, ctx context.Context) (domain.User, error)
}

type Router struct {
	useCase UseCase
}

func New(useCase UseCase) Router {
	return Router{
		useCase: useCase,
	}
}
