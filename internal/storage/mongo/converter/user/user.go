package user

import (
	"todoApp/internal/domain"
	"todoApp/internal/storage/mongo/object"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ToMongo(user domain.User) (object.User, error) {
	id, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return object.User{}, err
	}

	return object.User{
		ID:       id,
		Login:    user.Login,
		Password: user.Password,
		Email:    user.Email,
	}, nil
}

func ToDomain(user object.User) domain.User {
	return domain.User{
		ID:       user.ID.Hex(),
		Login:    user.Login,
		Password: user.Password,
		Email:    user.Email,
	}
}
