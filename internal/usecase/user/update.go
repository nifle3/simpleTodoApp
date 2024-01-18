package user

import (
	"context"
	"todoApp/pkg/hashing"
)

func (u UseCase) UpdatePassword(userID, password string, ctx context.Context) error {
	hashPassword, err := hashing.HashPassword(password)

	if err != nil {
		return err
	}

	return u.storage.UpdatePassword(hashPassword, userID, ctx)
}

func (u UseCase) UpdateEmail(userID, email string, ctx context.Context) error {
	return u.storage.UpdateEmail(email, userID, ctx)
}

func (u UseCase) UpdateLogin(userID, login string, ctx context.Context) error {
	return u.storage.UpdateLogin(login, userID, ctx)
}
