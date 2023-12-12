package controller

import (
	UserService "Plane-Booking/app/service/user"
	models "Plane-Booking/db/Models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddUser(context *gin.Context) {

	var input models.User

	bindingErr := context.ShouldBindJSON(&input)

	if bindingErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": true,
			"error":   "Expected a user type object in req body",
		})

		return
	}

	userObject, err := UserService.AddUser(&input)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": true,
			"error":   "Error in saving doc in DB",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"success": true,
		"user":    userObject,
	})

	fmt.Println("Running Add User Controller")
}

func GetUser(context *gin.Context) {

	var input = context.Query("id")

	user, err := UserService.GetUser(input)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": true,
			"error":   "Error in saving doc in DB",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"success": true,
		"user":    user,
	})
}
