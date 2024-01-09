package todo

import (
	"context"

	"todoApp/internal/domain"
)

func (u UseCase) AddTodo(userId string, todo domain.Todo, ctx context.Context) error {
	return u.storage.AddTodo(userId, todo, ctx)
}

func (u UseCase) DeleteTodo(userId, todoId string, ctx context.Context) error {
	return u.storage.DeleteTodo(userId, todoId, ctx)
}

func (u UseCase) UpdateTodo(userId string, todo domain.Todo, ctx context.Context) error {
	return u.storage.UpdateTodo(userId, todo, ctx)
}

func (u UseCase) GetAllTodo(userId string, ctx context.Context) ([]domain.Todo, error) {
	return u.storage.GetAllTodo(userId, ctx)
}
