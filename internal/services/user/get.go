package user

import (
	"context"
	"todoApp/internal/models"
)

func (u UseCase) Get(userID string, ctx context.Context) (models.User, error) {
	return u.storage.Get(userID, ctx)
}
