package main

import (
	"context"
	"time"
	"todoApp/internal/domain"
	"todoApp/internal/storage/mongo"
)

func main() {
	db, _ := mongo.NewStorage("mongodb://localhost:27017", context.TODO())

	todo := domain.Todo{
		Name:        "Магазин",
		DeadLine:    time.Now(),
		Description: "пойти в магаз",
		IsComplete:  false,
	}

	db.AddTodo("659cc9ed48ef867df238c19b", todo, context.TODO())

}
