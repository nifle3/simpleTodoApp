package todo

import (
	"context"
	"encoding/json"
	"net/http"
	"todoApp/internal/models"
)

// @Summary	Add todo for user
// @Tags		Todo
// @Accept		json
// @Produce	json
// @Param		input	body		models.Todo	false	"todo"
// @Failure	401		{object}	error
// @Failure	400		{object}	error
// @Failure	500		{object}	error
// @Success	201
// @Router		/v1/todo [put]
func (r Router) Add(userID string, w http.ResponseWriter, rq *http.Request) {
	var result models.Todo

	if err := json.NewDecoder(rq.Body).Decode(&result); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err := r.useCase.Add(userID, result, context.Background()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
