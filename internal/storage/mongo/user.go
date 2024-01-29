package mongo

import (
	"context"
	"todoApp/internal/models"
	converter "todoApp/internal/storage/mongo/converter/user"
	"todoApp/internal/storage/mongo/object"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s Storage) AddUser(user models.User, ctx context.Context) error {
	insertStruct, err := converter.ToMongo(user)
	if err != nil {
		return err
	}

	_, err = s.userCollection.InsertOne(ctx, insertStruct)

	return err
}

func (s Storage) UpdateLogin(login, userId string, ctx context.Context) error {
	return s.updateUserField(userId, "login", login, ctx)
}

func (s Storage) UpdatePassword(password, userId string, ctx context.Context) error {
	return s.updateUserField(userId, "password", password, ctx)
}

func (s Storage) UpdateEmail(email, userId string, ctx context.Context) error {
	return s.updateUserField(userId, "email", email, ctx)
}

func (s Storage) updateUserField(userId, key, value string, ctx context.Context) error {
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	updatedStruct := bson.D{{
		Key: "$set",
		Value: bson.M{
			key: value,
		},
	}}

	_, err = s.userCollection.UpdateOne(ctx, bson.M{"_id": id}, updatedStruct)

	return err
}

func (s Storage) DeleteUser(userId string, ctx context.Context) error {
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	_, err = s.userCollection.DeleteOne(ctx, bson.M{"_id": id})

	return err
}

func (s Storage) CheckUserExist(email string, ctx context.Context) (models.User, error) {
	var resultUser object.User

	err := s.userCollection.FindOne(ctx, bson.M{"email": email}).Decode(&resultUser)
	if err != nil {
		return models.User{}, err
	}

	return converter.ToDomain(resultUser), nil
}

func (s Storage) Get(userID string, ctx context.Context) (models.User, error) {
	filter := bson.D{{Key: "_id", Value: userID}}
	result := s.userCollection.FindOne(ctx, filter)
	if result.Err() != nil {
		return models.User{}, result.Err()
	}

	var user object.User
	if err := result.Decode(&user); err != nil {
		return models.User{}, err
	}

	return converter.ToDomain(user), nil
}
