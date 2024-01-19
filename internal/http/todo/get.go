package todo

import (
	"context"
	"encoding/json"
	"net/http"
)

const (
	todoQuery = "todo_id"
)

func (r Router) GetOne(userID string, w http.ResponseWriter, rq *http.Request) {
	w.Header().Add("Content-type", "application/json")

	if !rq.URL.Query().Has(todoQuery) {
		http.Error(w, "have not todo_id parametr", http.StatusBadRequest)
		return
	}

	todoID := rq.URL.Query().Get(todoQuery)

	result, err := r.useCase.GetOne(userID, todoID, context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
}

func (r Router) GetAll(userID string, w http.ResponseWriter, rq *http.Request) {
	w.Header().Add("Content-type", "application/json")

	result, err := r.useCase.GetAll(userID, context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
}
