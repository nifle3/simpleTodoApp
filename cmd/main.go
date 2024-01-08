package main

import (
	"context"
	"fmt"
	"todoApp/internal/domain"
	"todoApp/internal/storage"
)

func main() {
	a := "hello world"
	fmt.Printf("%s", a)

	user := domain.User{
		Login:    "qwe",
		Password: "qwe",
	}

	db, _ := storage.NewStorage("", context.TODO())

	db.AddUser(user, context.TODO())
}
