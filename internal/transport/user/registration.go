package user

import (
	"context"
	"encoding/json"
	"net/http"
	"todoApp/internal/models"
)

//	@Summary	Registration user
//	@Tags		Auth service
//	@Accept		json
//	@Produce	json
//	@Param		input	body		models.User	true	"new user"
//	@Failure	400		{object}	error
//	@Failure	401		{object}	error
//	@Failure	500		{object}	error
//	@Success	200
//	@Router		/v1/registration [put]
func (r Router) Registration(w http.ResponseWriter, rq *http.Request) {
	var result models.User

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
