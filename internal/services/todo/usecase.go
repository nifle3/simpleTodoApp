package todo

import (
	"context"

	"todoApp/internal/models"
)

type Storage interface {
	AddTodo(userID string, todo models.Todo, ctx context.Context) error
	DeleteTodo(userID, todoID string, ctx context.Context) error
	UpdateTodo(userID string, todo models.Todo, ctx context.Context) error
	GetAllTodo(userID string, ctx context.Context) ([]models.Todo, error)
	GetOneTodo(userID, todoID string, ctx context.Context) (models.Todo, error)
}

type UseCase struct {
	storage Storage
}

func NewUseCase(storage Storage) UseCase {
	return UseCase{
		storage: storage,
	}
}
