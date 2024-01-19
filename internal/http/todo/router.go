package todo

import (
	"context"
	"todoApp/internal/domain"
)

type UseCase interface {
	Add(userID string, todo domain.Todo, ctx context.Context) error
	Delete(userID, todoId string, ctx context.Context) error
	Update(userID string, todo domain.Todo, ctx context.Context) error
	GetOne(userID, todoID string, ctx context.Context) (domain.Todo, error)
	GetAll(userID string, ctx context.Context) ([]domain.Todo, error)
}

type Router struct {
	useCase UseCase
}

func New(useCase UseCase) Router {
	return Router{
		useCase: useCase,
	}
}
