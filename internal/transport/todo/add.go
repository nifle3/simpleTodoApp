package todo

import (
	"context"
	"encoding/json"
	"net/http"
	"todoApp/internal/models"
)

func (r Router) Add(userID string, w http.ResponseWriter, rq *http.Request) {
	var result models.Todo

	if err := json.NewDecoder(rq.Body).Decode(&result); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err := r.useCase.Add(userID, result, context.Background()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
