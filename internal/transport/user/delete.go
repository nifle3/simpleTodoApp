package user

import (
	"context"
	"net/http"
)

// @Summary	Delete user
// @Tags		User
// @Accept		json
// @Produce	json
// @Security ApiKeyAuth
// @Failure	400	{object}	error
// @Failure	401	{object}	error
// @Failure	500	{object}	error
// @Success	200
// @Router		/v1/user [delete]
func (r Router) Delete(userID string, w http.ResponseWriter, rq *http.Request) {
	if err := r.useCase.Delete(userID, context.Background()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
