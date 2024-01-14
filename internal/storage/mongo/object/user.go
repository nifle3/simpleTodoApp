package object

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Login    string             `bson:"login"`
	Password string             `bson:"password"`
	Email    string             `bson:"email"`
}
