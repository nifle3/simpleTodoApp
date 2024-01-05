package storage

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbName  = "TodoDB"
	colName = "users"
)

type Storage struct {
	db mongo.Client
}

func NewStorage(uri string) (*Storage, error) {
	db, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))

	if err != nil {
		return nil, fmt.Errorf("can't connect to database")
	}

	defer db.Disconnect(context.Background())

	if err = db.Ping(context.Background(), nil); err != nil {
		return nil, fmt.Errorf("Cannot ping to database")
	}

	return &Storage{
		db: *db,
	}, nil
}
