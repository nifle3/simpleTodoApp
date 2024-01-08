package storage

import "todoApp/internal/domain"

func (s Storage) AddTodo(todo domain.Todo) error {
	return nil
}

func (s Storage) DeleteTodo(todo domain.Todo) error {
	return nil
}

func (s Storage) UpdateTodo(todo domain.Todo) error {
	return nil
}

func GetAllTodo(user domain.User) ([]domain.Todo, error) {
	return nil
}
