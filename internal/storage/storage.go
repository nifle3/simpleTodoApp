package storage

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TODO: add uri path
const uri string = ""

type Storage struct {
	db *mongo.Client
}

// FIXME: change disconnect
// TODO: add new options to client
func NewStorage() (*Storage, error) {
	db, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		return nil, fmt.Errorf("can't connect to database")
	}

	defer db.Disconnect(context.Background())

	if err = db.Ping(context.TODO(), nil); err != nil {
		return nil, fmt.Errorf("Cannot ping to database")
	}

	return &Storage{
		db: db,
	}, nil
}
