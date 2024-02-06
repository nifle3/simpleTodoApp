package user

import (
	"time"
	"todoApp/internal/models"
	"todoApp/internal/storage/mongo/object"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ToMongoWithNewID(user models.User) object.User {
	id := primitive.NewObjectIDFromTimestamp(time.Now())

	return object.User{
		ID:       id,
		Login:    user.Login,
		Password: user.Password,
		Email:    user.Email,
	}
}

func ToDomain(user object.User) models.User {
	return models.User{
		Login:    user.Login,
		Password: user.Password,
		Email:    user.Email,
	}
}
