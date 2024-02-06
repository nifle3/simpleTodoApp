package mongo

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
	db *mongo.Client
}

func New(uri string, ctx context.Context) (*Storage, error) {
	db, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		return nil, fmt.Errorf("can't connect to database")
	}

	if err = db.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("Cannot ping to database")
	}

	return &Storage{
		db: db,
	}, nil
}

func (s Storage) GetCollections() *mongo.Collection {
	return s.db.Database(dbName).Collection(colName)
}
