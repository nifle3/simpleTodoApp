package todo

import (
	"context"

	"todoApp/internal/domain"
)

func (u UseCase) Add(userID string, todo domain.Todo, ctx context.Context) error {
	return u.storage.AddTodo(userID, todo, ctx)
}
