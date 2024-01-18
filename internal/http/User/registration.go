package user

import (
	"context"
	"encoding/json"
	"net/http"
	"todoApp/internal/domain"
)

func (r Router) Registration(w http.ResponseWriter, rq *http.Request) {
	var result domain.User

	if err := json.NewDecoder(rq.Body).Decode(&result); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := r.useCase.Add(result, context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
