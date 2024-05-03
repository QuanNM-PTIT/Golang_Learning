package routes

import "github.com/gin-gonic/gin"

func RegisterServer(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventById)
	server.POST("/events/update-event/:id", UpdateEvent)
	server.POST("/events/delete-event/:id", DeleteEvent)
	server.POST("/events/create-event", createEvent)
	server.POST("/users/create-user", CreateUser)
	server.POST("/users/login", Login)
}
