package todo

import (
	"context"
	"todoApp/internal/domain"
)

func (u UseCase) GetAll(userID string, ctx context.Context) ([]domain.Todo, error) {
	return u.storage.GetAllTodo(userID, ctx)
}

func (u UseCase) GetOne(userID, todoID string, ctx context.Context) (domain.Todo, error) {
	return u.storage.GetOneTodo(userID, todoID, ctx)
}
