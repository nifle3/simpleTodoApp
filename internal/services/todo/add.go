package todo

import (
	"context"

	"todoApp/internal/models"
)

func (u UseCase) Add(userID string, todo models.Todo, ctx context.Context) error {
	return u.storage.AddTodo(userID, todo, ctx)
}
