package user

import (
	"context"

	"todoApp/internal/models"
	"todoApp/pkg/hashing"
)

func (u UseCase) Add(user models.User, ctx context.Context) error {
	password, err := hashing.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = password

	return u.storage.AddUser(user, ctx)
}
