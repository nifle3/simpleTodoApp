package user

import (
	"context"
	"encoding/json"
	"net/http"
)

// @Summary	Get user
// @Tags		User
// @Accept		json
// @Produce	json
// @Security ApiKeyAuth
// @Failure	400	{object}	error
// @Failure	401	{object}	error
// @Failure	500	{object}	error
// @Success	200
// @Router		/v1/user [get]
func (r Router) Get(userID string, w http.ResponseWriter, rq *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	result, err := r.useCase.Get(userID, context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
