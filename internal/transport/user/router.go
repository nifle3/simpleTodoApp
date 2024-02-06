package user

import (
	"context"
	"todoApp/internal/models"
)

type UseCase interface {
	Login(email, password string, ctx context.Context) (string, error)
	Add(user models.User, ctx context.Context) error
	Delete(userID string, ctx context.Context) error
	UpdatePassword(userID, password string, ctx context.Context) error
	UpdateEmail(userID, email string, ctx context.Context) error
	UpdateLogin(userID, login string, ctx context.Context) error
	Get(userID string, ctx context.Context) (models.User, error)
}

type Router struct {
	useCase UseCase
}

func New(useCase UseCase) Router {
	return Router{
		useCase: useCase,
	}
}
