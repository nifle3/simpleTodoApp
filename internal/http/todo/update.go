package todo

import (
	"encoding/json"
	"net/http"
	"todoApp/internal/domain"
)

func (r Router) Update(userID string, w http.ResponseWriter, rq *http.Request) error {
	var result domain.Todo

	if err := json.NewDecoder(rq.Body).Decode(&result); err != nil {
		return err
	}

	return nil
}
