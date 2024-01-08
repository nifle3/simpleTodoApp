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
	userCollection *mongo.Collection
}

func NewStorage(uri string, ctx context.Context) (*Storage, error) {
	db, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		return nil, fmt.Errorf("can't connect to database")
	}

	defer db.Disconnect(ctx)

	if err = db.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("Cannot ping to database")
	}

	collection := db.Database(dbName).Collection(colName)

	return &Storage{
		userCollection: collection,
	}, nil
}
