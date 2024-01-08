package storage

import (
	"context"
	"todoApp/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
)

func (s Storage) AddUser(user domain.User, ctx context.Context) error {
	_, err := s.userCollection.InsertOne(ctx, user)

	return err
}

func (s Storage) UpdateUser(user domain.User, ctx context.Context) error {
	_, err := s.userCollection.UpdateByID(ctx, user.ID, user)

	return err
}

func (s Storage) DeleteUser(user domain.User, ctx context.Context) error {
	_, err := s.userCollection.DeleteOne(ctx, bson.D{{"_id", user.ID}})

	return err
}

func (s Storage) CheckUserExist(user domain.User, ctx context.Context) (domain.User, error) {
	result := s.userCollection.FindOne(ctx, bson.D{{"_id", user.ID}})
	var resultUser domain.User
	if result.Err() != nil {
		return resultUser, result.Err()
	}

	result.Decode(&resultUser)
	return resultUser, nil
}
