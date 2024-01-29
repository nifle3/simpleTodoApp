package user

import (
	"context"
	"todoApp/internal/models"
	"todoApp/pkg/hashing"
)

func (u UseCase) Login(email, password string, ctx context.Context) (models.User, error) {
	user, err := u.storage.CheckUserExist(email, ctx)
	if err != nil {
		return user, err
	}

	if err := hashing.CheckHashPassword(user.Password, password); err != nil {
		return user, err
	}

	return user, nil
}
