package user

import (
	"context"
	"net/http"
)

func (r Router) Delete(userID string, w http.ResponseWriter, rq *http.Request) error {
	return r.useCase.Delete(userID, context.Background())
}
