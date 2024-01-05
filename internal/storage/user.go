package storage

import (
	"context"
	"todoApp/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
)

func (s Storage) AddUser(user domain.User) error {
	col := s.db.Database(dbName).Collection(colName)
	_, err := col.InsertOne(context.Background(), user)

	return err
}

func (s Storage) UpdateUser(user domain.User) error {
	col := s.db.Database(dbName).Collection(colName)
	_, err := col.UpdateByID(context.Background(), user.ID, user)

	return err
}

func (s Storage) DeleteUser(user domain.User) error {
	col := s.db.Database(dbName).Collection(colName)
	_, err := col.DeleteOne(context.Background(), bson.D{{"_id", user.ID}})

	return err
}

func (s Storage) CheckUserExist(user domain.User) error {
	col := s.db.Database(dbName).Collection(colName)
	result, err := col.FindOne(context.Background(), bson.D{{"_id", user.ID}})
	return nil
}
