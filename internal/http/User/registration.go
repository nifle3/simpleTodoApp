package user

import (
	"context"
	"encoding/json"
	"net/http"
	"todoApp/internal/domain"
)

func (r Router) Registration(w http.ResponseWriter, rq *http.Request) error {
	var result domain.User

	if err := json.NewDecoder(rq.Body).Decode(&result); err != nil {
		return err
	}

	ctx := context.Background()
	err := r.useCase.Add(result, ctx)
	if err != nil {
		return err
	}

	return nil
}
