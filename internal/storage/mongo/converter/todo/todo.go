package todo

import (
	"todoApp/internal/models"
	"todoApp/internal/storage/mongo/object"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ToMongo(todo models.Todo) (object.Todo, error) {
	id, err := primitive.ObjectIDFromHex(todo.ID)
	if err != nil {
		return object.Todo{}, err
	}

	return object.Todo{
		ID:          id,
		Name:        todo.Name,
		Description: todo.Description,
		Deadline:    todo.DeadLine,
		IsComplete:  todo.IsComplete,
	}, nil
}

func ToModel(todo object.Todo) models.Todo {
	return models.Todo{
		ID:          todo.ID.Hex(),
		Name:        todo.Name,
		DeadLine:    todo.Deadline,
		Description: todo.Description,
		IsComplete:  todo.IsComplete,
	}
}
