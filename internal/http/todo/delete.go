package todo

import (
	"context"
	"net/http"
)

func (r Router) Delete(userID string, w http.ResponseWriter, rq *http.Request) error {
	todoID := rq.PostFormValue("todo_id")

	if err := r.useCase.DeleteTodo(userID, todoID, context.Background()); err != nil {
		return err
	}

	return nil
}
