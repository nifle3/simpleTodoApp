package user

import (
	"context"
	"todoApp/internal/domain"
)

type Storage interface {
	AddUser(user domain.User, ctx context.Context) error
	UpdateUser(user domain.User, ctx context.Context) error
	DeleteUser(userId string, ctx context.Context) error
	CheckUserExist(email string, ctx context.Context) (domain.User, error)
}

type UseCase struct {
	storage Storage
}

func NewUseCase(storage Storage) UseCase {
	return UseCase{
		storage: storage,
	}
}
