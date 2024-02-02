package todo

import (
	"context"
	"encoding/json"
	"net/http"
	"todoApp/internal/models"
)

// @Summary	Update todo for user
// @Tags		Todo
// @Accept		json
// @Produce	json
// @Security ApiKeyAuth
// @Param		input	body		models.Todo	true	"used for updated created todo"
// @Failure	400		{object}	error
// @Failure	401		{object}	error
// @Failure	500		{object}	error
// @Success	200
// @Router		/v1/todo [patch]
func (r Router) Update(userID string, w http.ResponseWriter, rq *http.Request) {
	var result models.Todo

	if err := json.NewDecoder(rq.Body).Decode(&result); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := r.useCase.Update(userID, result, context.Background()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
