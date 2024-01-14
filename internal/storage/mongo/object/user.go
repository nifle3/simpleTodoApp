package object

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	Login    string             `bson:"login"`
	Password string             `bson:"password"`
	Email    string             `bson:"email"`
}
