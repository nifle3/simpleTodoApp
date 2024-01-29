package user

import (
	"context"
	"net/http"
)

func (r Router) Delete(userID string, w http.ResponseWriter, rq *http.Request) {
	if err := r.useCase.Delete(userID, context.Background()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
