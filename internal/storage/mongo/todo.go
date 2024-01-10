package mongo

import (
	"context"
	"net/http"
	"todoApp/internal/domain"
	"todoApp/pkg/httpError"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s Storage) AddTodo(userId string, todo domain.Todo, ctx context.Context) error {
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return httpError.Error{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	todo.ID = primitive.NewObjectID().Hex()

	push := bson.D{{Key: "$push", Value: bson.D{{Key: "todos", Value: todo}}}}

	_, err = s.userCollection.UpdateByID(ctx, id, push)

	if err != nil {
		return httpError.Error{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return httpError.Error{
		StatusCode: http.StatusCreated,
		Message:    "Todo created",
	}
}

func (s Storage) DeleteTodo(userId string, todoId string, ctx context.Context) httpError.Error {
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return httpError.Error{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	pull := bson.D{{Key: "$pull", Value: bson.D{{Key: "todos", Value: bson.D{{Key: "id", Value: todoId}}}}}}
	if _, err = s.userCollection.UpdateByID(ctx, id, pull); err != nil {
		return httpError.Error{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return httpError.Error{
		StatusCode: http.StatusOK,
		Message:    "todo deleted",
	}
}

func (s Storage) UpdateTodo(userId string, todo domain.Todo, ctx context.Context) httpError.Error {
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return httpError.Error{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	filter := bson.D{{Key: "_id", Value: id}, {Key: "todos.id", Value: todo.ID}}
	updated := bson.D{{Key: "$set", Value: bson.M{
		"todos.$.name":        todo.Name,
		"todos.$.deadline":    todo.DeadLine,
		"todos.$.description": todo.Description,
		"todo.$.iscomplete":   todo.IsComplete,
	}}}

	_, err = s.userCollection.UpdateOne(ctx, filter, updated)

	if err != nil {
		return httpError.Error{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return httpError.Error{
		StatusCode: http.StatusOK,
		Message:    "todo updated",
	}
}

func (s Storage) GetAllTodo(userId string, ctx context.Context) ([]domain.Todo, httpError.Error) {
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, httpError.Error{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	var todos struct {
		Todos []domain.Todo `bson:"todos"`
	}

	if err = s.userCollection.FindOne(ctx, bson.D{{Key: "_id", Value: id}}).Decode(&todos); err != nil {
		return nil, httpError.Error{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return todos.Todos, httpError.Error{
		StatusCode: http.StatusOK,
		Message:    "get all todos",
	}
}
