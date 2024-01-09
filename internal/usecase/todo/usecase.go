package todo

import (
	"context"

	"todoApp/internal/domain"
)

type Storage interface {
	AddTodo(userId string, todo domain.Todo, ctx context.Context) error
	DeleteTodo(userId string, todoId string, ctx context.Context) error
	UpdateTodo(userId string, todo domain.Todo, ctx context.Context) error
	GetAllTodo(userId string, ctx context.Context) ([]domain.Todo, error)
}

type UseCase struct {
	storage Storage
}

func NewUseCase(storage Storage) UseCase {
	return UseCase{
		storage: storage,
	}
}
