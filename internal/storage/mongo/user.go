package mongo

import (
	"context"
	"todoApp/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s Storage) AddUser(user domain.User, ctx context.Context) error {
	insertStruct := bson.M{
		"_id":      primitive.NewObjectID(),
		"login":    user.Login,
		"password": user.Password,
		"email":    user.Email,
		"todos":    make([]domain.Todo, 0),
	}

	_, err := s.userCollection.InsertOne(ctx, insertStruct)

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

func (s Storage) CheckUserExist(email string, ctx context.Context) (domain.User, error) {
	var resultUser struct {
		Id       primitive.ObjectID `bson:"_id"`
		Email    string             `bson:"email"`
		Password string             `bson:"password"`
		Login    string             `bson:"login"`
	}

	err := s.userCollection.FindOne(ctx, bson.M{"email": email}).Decode(&resultUser)

	user := domain.User{
		ID:       resultUser.Id.Hex(),
		Email:    resultUser.Email,
		Password: resultUser.Password,
		Login:    resultUser.Login,
	}

	return user, err
}
