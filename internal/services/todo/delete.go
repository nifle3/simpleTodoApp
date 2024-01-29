package todo

import "context"

func (u UseCase) Delete(userID, todoID string, ctx context.Context) error {
	return u.storage.DeleteTodo(userID, todoID, ctx)
}
