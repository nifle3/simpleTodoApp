package user

import (
	"context"

	"todoApp/internal/domain"
	"todoApp/pkg/hashing"
)

func (u UseCase) Add(user domain.User, ctx context.Context) error {
	password, err := hashing.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = password

	return u.storage.AddUser(user, ctx)
}

func (u UseCase) Delete(user domain.User, ctx context.Context) error {
	return u.storage.DeleteUser(user.ID, ctx)
}

func (u UseCase) UpdatePassword(userId, password string, ctx context.Context) error {
	hashPassword, err := hashing.HashPassword(password)

	if err != nil {
		return err
	}

	return u.storage.UpdatePassword(hashPassword, userId, ctx)
}

func (u UseCase) UpdateEmail(userId, email string, ctx context.Context) error {
	return u.storage.UpdateEmail(email, userId, ctx)
}

func (u UseCase) UpdateLogin(userId, login string, ctx context.Context) error {
	return u.storage.UpdateLogin(login, userId, ctx)
}

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
