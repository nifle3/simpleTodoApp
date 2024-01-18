package user

import (
	"context"
	"todoApp/internal/domain"
	"todoApp/pkg/hashing"
)

func (u UseCase) Login(email, password string, ctx context.Context) (domain.User, error) {
	user, err := u.storage.CheckUserExist(email, ctx)
	if err != nil {
		return user, err
	}

	if err := hashing.CheckHashPassword(user.Password, password); err != nil {
		return user, err
	}

	return user, nil
}
