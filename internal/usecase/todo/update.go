package todo

import (
	"context"
	"todoApp/internal/domain"
)

func (u UseCase) Update(userID string, todo domain.Todo, ctx context.Context) error {
	return u.storage.UpdateTodo(userID, todo, ctx)
}
