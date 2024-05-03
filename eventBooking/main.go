package main

import (
	"eventBooking/database"
	"eventBooking/models"
	"eventBooking/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	server := gin.Default()
	err := godotenv.Load()

	if err != nil {
		panic("Failed to load .env file!")
		return
	}

	err = database.InitDB()

	if err != nil {
		panic("Failed to connect to database!")
		return
	}

	models.Migrate()

	routes.RegisterServer(server)

	err = server.Run(":8080")

	if err != nil {
		panic("Failed to start server!")
	}
}
