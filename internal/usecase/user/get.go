package user

import (
	"context"
	"todoApp/internal/domain"
)

func (u UseCase) Get(userID string, ctx context.Context) (domain.User, error) {
	return u.storage.Get(userID, ctx)
}
