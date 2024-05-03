package routes

import (
	"eventBooking/models"
	"eventBooking/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request!"})
		return
	}

	user.Password, err = utils.HashPassword(user.Password)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sign up!"})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sign up!"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"status": "User created successfully"})
}

func Login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request!"})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials!"})
		return
	}

	var accessToken string
	accessToken, err = utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login!"})
		fmt.Println(err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}
