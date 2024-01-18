package todo

import (
	"context"

	"todoApp/internal/domain"
)

type Storage interface {
	AddTodo(userID string, todo domain.Todo, ctx context.Context) error
	DeleteTodo(userID, todoID string, ctx context.Context) error
	UpdateTodo(userID string, todo domain.Todo, ctx context.Context) error
	GetAllTodo(userID string, ctx context.Context) ([]domain.Todo, error)
	GetOneTodo(userID, todoID string, ctx context.Context) (domain.Todo, error)
}

type UseCase struct {
	storage Storage
}

func NewUseCase(storage Storage) UseCase {
	return UseCase{
		storage: storage,
	}
}
