package todo

import (
	"context"
	"todoApp/internal/models"
)

type UseCase interface {
	Add(userID string, todo models.Todo, ctx context.Context) error
	Delete(userID, todoId string, ctx context.Context) error
	Update(userID string, todo models.Todo, ctx context.Context) error
	GetOne(userID, todoID string, ctx context.Context) (models.Todo, error)
	GetAll(userID string, ctx context.Context) ([]models.Todo, error)
}

type Router struct {
	useCase UseCase
}

func New(useCase UseCase) Router {
	return Router{
		useCase: useCase,
	}
}
