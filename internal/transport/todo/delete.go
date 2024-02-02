package todo

import (
	"context"
	"net/http"
)

// @Summary	Delete todo from user
// @Tags		Todo
// @Accept		json
// @Produce	json
// @Security ApiKeyAuth
// @Param		todo_id	formData	string	true	"todo id in stirng"
// @Failure	400		{object}	error
// @Failure	401		{object}	error
// @Success	200
// @Router		/v1/todo [delete]
func (r Router) Delete(userID string, w http.ResponseWriter, rq *http.Request) {
	todoID := rq.PostFormValue("todo_id")

	if err := r.useCase.Delete(userID, todoID, context.Background()); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
