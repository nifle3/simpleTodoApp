package todo

import (
	"context"
	"encoding/json"
	"net/http"
	"todoApp/internal/domain"
)

func (r Router) Add(w http.ResponseWriter, rq *http.Request) {
	var result domain.Todo

	if err := json.NewDecoder(rq.Body).Decode(&result); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err := r.useCase.AddTodo("", result, context.Background()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
