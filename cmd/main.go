package main

import (
	"log"
	"todoApp/internal/app"

	"github.com/joho/godotenv"
)

//	@title			todo app
//	@version		1.0
//	@description	simple golang pet project
//	@termsOfService	http://swagger.io/terms/

//	@produce	json
//	@accept		json

//  @securitydefinitions.apikey ApiKeyAuth
//  @in header
//  @name jwt token

//	@license.name	MIT

//	@host		:8080

// @scheme	http
func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("No .env file found")
	}

	app.Start()
}
