package user

import (
	"context"
	"todoApp/pkg/hashing"
)

func (u UseCase) Login(email, password string, ctx context.Context) (string, error) {
	password, userID, err := u.storage.CheckUserExist(email, ctx)
	if err != nil {
		return "", err
	}

	if err := hashing.CheckHashPassword(password, password); err != nil {
		return "", err
	}

	return userID, nil
}
