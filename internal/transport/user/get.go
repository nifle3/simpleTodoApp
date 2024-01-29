package user

import (
	"context"
	"encoding/json"
	"net/http"
)

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
