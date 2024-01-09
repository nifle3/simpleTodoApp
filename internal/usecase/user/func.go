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

func (u UseCase) Update(user domain.User, ctx context.Context) error {

}
