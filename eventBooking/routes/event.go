package routes

import (
	"eventBooking/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getEvents(context *gin.Context) {
	events, err := models.GetEvents()
	if err != nil {
		context.JSON(404, gin.H{"error": "No events found!"})
		return
	}
	context.JSON(200, events)
}

func getEventById(context *gin.Context) {
	id := context.Param("id")
	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(404, gin.H{"error": "Event not found!"})
		return
	}
	context.JSON(200, event)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save event!"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"status": "Event created successfully"})
}

func UpdateEvent(context *gin.Context) {
	id := context.Param("id")
	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(404, gin.H{"error": "Event not found!"})
		return
	}
	err = context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request!"})
		return
	}
	err = event.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save event!"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"status": "Event updated successfully"})
}

func DeleteEvent(context *gin.Context) {
	id := context.Param("id")
	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(404, gin.H{"error": "Event not found!"})
		return
	}
	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete event!"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"status": "Event deleted successfully"})
}
