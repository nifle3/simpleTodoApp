package todo

import (
	"context"
	"todoApp/internal/domain"
)

type UseCase interface {
	AddTodo(userId string, todo domain.Todo, ctx context.Context) error
	DeleteTodo(userId, todoId string, ctx context.Context) error
	UpdateTodo(userId string, todo domain.Todo, ctx context.Context) error
}

type Router struct {
	useCase UseCase
}
