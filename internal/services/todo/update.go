package todo

import (
	"context"
	"todoApp/internal/models"
)

func (u UseCase) Update(userID string, todo models.Todo, ctx context.Context) error {
	return u.storage.UpdateTodo(userID, todo, ctx)
}
