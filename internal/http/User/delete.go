package user

import (
	"context"
	"net/http"
)

func (r Router) Delete(w http.ResponseWriter, rq *http.Request) {
	if err := r.useCase.Delete("", context.Background()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
