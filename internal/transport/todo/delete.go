package todo

import (
	"context"
	"net/http"
)

func (r Router) Delete(userID string, w http.ResponseWriter, rq *http.Request) {
	todoID := rq.PostFormValue("todo_id")

	if err := r.useCase.Delete(userID, todoID, context.Background()); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
