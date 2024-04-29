package main

import (
	"eventBooking/database"
	"eventBooking/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
)

func main() {
	server := gin.Default()
	err := godotenv.Load()

	if err != nil {
		panic("Failed to load .env file!")
		return
	}

	database.InitDB()
	models.Migrate()

	server.GET("/", getEvents)
	server.POST("/create-event", createEvent)

	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events := models.GetEvents()
	context.JSON(200, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	event.Save()
	context.JSON(http.StatusCreated, gin.H{"status": "Event created successfully"})
}
