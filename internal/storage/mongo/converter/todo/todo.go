package todo

import (
	"todoApp/internal/domain"
	"todoApp/internal/storage/mongo/object"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ToMongo(todo domain.Todo) (object.Todo, error) {
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

func ToDomain(todo object.Todo) domain.Todo {
	return domain.Todo{
		ID:          todo.ID.Hex(),
		Name:        todo.Name,
		DeadLine:    todo.Deadline,
		Description: todo.Description,
		IsComplete:  todo.IsComplete,
	}
}
