package todo

import (
	"context"
	"encoding/json"
	"net/http"
)

const (
	todoQuery = "todo_id"
)

//		@Summary	Get todo
//		@Tags		Todo
//		@Accept		json
//		@Produce	json
//	 @Security ApiKeyAuth
//		@Param		todo_id	query		string	true	"todo id which get"
//		@Failure	400		{object}	error
//		@Failure	401		{object}	error
//		@Failure	500		{object}	error
//		@Success	200		{object}	models.Todo	"one model"
//		@Router		/v1/todo/{id} [get]
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

	w.WriteHeader(http.StatusOK)
}

//	@Summary	Get all user's todo
//	@Tags		Todo
//	@Accept		json
//	@Produce	json
//
// @Security ApiKeyAuth
//
//	@Failure	400	{object}	error
//	@Failure	401	{object}	error
//	@Failure	500	{object}	error
//	@Success	200	{object}	[]models.Todo	"many models"
//	@Router		/v1/get [get]
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

	w.WriteHeader(http.StatusOK)
}
