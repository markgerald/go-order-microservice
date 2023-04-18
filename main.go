package main

import (
	"github.com/joho/godotenv"
	"github.com/markgerald/go-order-microservice/database"
	"github.com/markgerald/go-order-microservice/routes"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.DbConnect()
	routes.HandleRequests()
}
