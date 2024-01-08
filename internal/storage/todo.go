package storage

import (
	"context"
	"todoApp/internal/domain"
)

func (s Storage) AddTodo(todo domain.Todo, ctx context.Context) error {
	return nil
}

func (s Storage) DeleteTodo(todo domain.Todo, ctx context.Context) error {
	return nil
}

func (s Storage) UpdateTodo(todo domain.Todo, ctx context.Context) error {
	return nil
}

func (s Storage) GetAllTodo(user domain.User, ctx context.Context) ([]domain.Todo, error) {
	return nil, nil
}
