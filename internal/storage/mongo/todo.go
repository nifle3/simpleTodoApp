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

	todo.ID = primitive.NewObjectID().Hex()

	push := bson.D{{Key: "$push", Value: bson.D{{Key: "todos", Value: todo}}}}

	_, err = s.userCollection.UpdateByID(ctx, id, push)

	return err
}

func (s Storage) DeleteTodo(userId string, todoId string, ctx context.Context) httpError.Error {
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	pull := bson.D{{Key: "$pull", Value: bson.D{{Key: "todos", Value: bson.D{{Key: "id", Value: todoId}}}}}}
	_, err = s.userCollection.UpdateByID(ctx, id, pull)

	return err
}

func (s Storage) UpdateTodo(userId string, todo domain.Todo, ctx context.Context) httpError.Error {
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	filter := bson.D{{Key: "_id", Value: id}, {Key: "todos.id", Value: todo.ID}}
	updated := bson.D{{Key: "$set", Value: bson.M{
		"todos.$.name":        todo.Name,
		"todos.$.deadline":    todo.DeadLine,
		"todos.$.description": todo.Description,
		"todo.$.iscomplete":   todo.IsComplete,
	}}}

	_, err = s.userCollection.UpdateOne(ctx, filter, updated)

	return err
}

func (s Storage) GetAllTodo(userId string, ctx context.Context) ([]domain.Todo, error) {
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	var todos struct {
		Todos []domain.Todo `bson:"todos"`
	}

	if err = s.userCollection.FindOne(ctx, bson.D{{Key: "_id", Value: id}}).Decode(&todos); err != nil {
		return nil, err
	}

	return todos.Todos, nil
}
