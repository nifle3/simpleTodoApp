package main

import (
	"log"
	"todoApp/internal/app"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("No .env file found")
	}

	app.Start()
}
