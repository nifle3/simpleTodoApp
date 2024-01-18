package user

import (
	"context"
	"todoApp/internal/domain"
)

type Storage interface {
	AddUser(user domain.User, ctx context.Context) error
	UpdatePassword(password, userID string, ctx context.Context) error
	UpdateLogin(login, userID string, ctx context.Context) error
	UpdateEmail(email, userID string, ctx context.Context) error
	DeleteUser(userID string, ctx context.Context) error
	CheckUserExist(email string, ctx context.Context) (domain.User, error)
	Get(userID string, ctx context.Context) (domain.User, error)
}

type UseCase struct {
	storage Storage
}

func NewUseCase(storage Storage) UseCase {
	return UseCase{
		storage: storage,
	}
}
