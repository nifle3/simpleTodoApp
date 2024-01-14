package object

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID          primitive.ObjectID `bson:"id"`
	Name        string             `bson:"name"`
	Deadline    time.Time          `bson:"deadline"`
	Description string             `bson:"description"`
	IsComplete  bool               `bson:"is_complete"`
}
