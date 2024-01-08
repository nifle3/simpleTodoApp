package main

import (
	"context"
	"fmt"
	"todoApp/internal/domain"
	"todoApp/internal/storage"
)

func main() {
	a := "hello world"
	fmt.Printf("%s\n", a)

	user := domain.User{
		Login:    "qwe",
		Password: "qwe",
	}

	db, _ := storage.NewStorage("mongodb://127.0.0.1:27017", context.TODO())

	db.AddUser(user, context.TODO())
}
