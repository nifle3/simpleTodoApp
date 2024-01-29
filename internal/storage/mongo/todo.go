package mongo

import (
	"context"

	"todoApp/internal/models"
	converter "todoApp/internal/storage/mongo/converter/todo"
	"todoApp/internal/storage/mongo/object"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s Storage) AddTodo(userId string, todo models.Todo, ctx context.Context) error {
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	mongoTodo, err := converter.ToMongo(todo)

	push := bson.D{{Key: "$push", Value: bson.D{{Key: "todos", Value: mongoTodo}}}}

	_, err = s.userCollection.UpdateByID(ctx, id, push)

	return err
}

func (s Storage) DeleteTodo(userId string, todoId string, ctx context.Context) error {
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	todoObjectId, err := primitive.ObjectIDFromHex(todoId)

	pull := bson.D{{Key: "$pull", Value: bson.D{{Key: "todos", Value: bson.D{{Key: "id", Value: todoObjectId}}}}}}
	_, err = s.userCollection.UpdateByID(ctx, id, pull)

	return err
}

func (s Storage) UpdateTodo(userId string, todo models.Todo, ctx context.Context) error {
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	updatedObject, err := converter.ToMongo(todo)
	if err != nil {
		return err
	}

	filter := bson.D{{Key: "_id", Value: id}, {Key: "todos.id", Value: updatedObject.ID}}
	updated := bson.D{{Key: "$set", Value: updatedObject}}

	_, err = s.userCollection.UpdateOne(ctx, filter, updated)

	return err
}

func (s Storage) GetAllTodo(userId string, ctx context.Context) ([]models.Todo, error) {
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	var todos struct {
		Todos []object.Todo `bson:"todos"`
	}

	if err = s.userCollection.FindOne(ctx, bson.D{{Key: "_id", Value: id}}).Decode(&todos); err != nil {
		return nil, err
	}

	result := make([]models.Todo, len(todos.Todos))

	for idx := range result {
		result[idx] = converter.ToModel(todos.Todos[idx])
	}

	return result, nil
}

func (s Storage) GetOneTodo(userID, todoID string, ctx context.Context) (models.Todo, error) {
	userIDHex, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return models.Todo{}, err
	}

	todoIDHex, err := primitive.ObjectIDFromHex(todoID)
	if err != nil {
		return models.Todo{}, err
	}

	filter := bson.D{{Key: "_id", Value: userIDHex}, {Key: "todos.id", Value: todoIDHex}}

	result := s.userCollection.FindOne(ctx, filter)
	if result.Err() != nil {
		return models.Todo{}, result.Err()
	}

	var todo object.Todo
	if err := result.Decode(&todo); err != nil {
		return models.Todo{}, err
	}

	return converter.ToModel(todo), nil
}
