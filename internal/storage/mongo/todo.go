package mongo

import (
	"context"
	"todoApp/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s Storage) AddTodo(userId string, todo domain.Todo, ctx context.Context) error {
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	push := bson.D{{Key: "$push", Value: bson.D{{Key: "todos", Value: todo}}}}

	_, err = s.userCollection.UpdateByID(ctx, id, push)

	return err
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
