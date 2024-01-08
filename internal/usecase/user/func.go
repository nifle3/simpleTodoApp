package user

import (
	"context"
	"todoApp/internal/domain"
)

func (u UseCase) Add(user domain.User, ctx context.Context) error {
	if err := u.storage.AddUser(user, ctx); err != nil {
		return err
	}

	return nil
}
