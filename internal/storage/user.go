package storage

import (
	"context"
	"todoApp/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s Storage) AddUser(user domain.User, ctx context.Context) error {
	insertStruct := bson.D{
		{"_id", primitive.NewObjectID()},
		{"login", user.Login},
		{"password", user.Password},
	}

	_, err := s.userCollection.InsertOne(ctx, insertStruct)

	return err
}

func (s Storage) UpdateUser(user domain.User, ctx context.Context) error {
	id, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return err
	}

	updatedStruct := bson.D{
		{"login", user.Login},
		{"password", user.Password},
	}

	_, err = s.userCollection.UpdateByID(ctx, id, updatedStruct)

	return err
}

func (s Storage) DeleteUser(user domain.User, ctx context.Context) error {
	id, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return err
	}

	_, err = s.userCollection.DeleteOne(ctx, bson.D{{"_id", id}})

	return err
}

func (s Storage) CheckUserExist(user domain.User, ctx context.Context) (domain.User, error) {
	var resultUser domain.User

	id, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return resultUser, err
	}

	err = s.userCollection.FindOne(ctx, bson.D{{"_id", id}}).Decode(&resultUser)

	return resultUser, err
}
