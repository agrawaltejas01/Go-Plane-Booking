package controller

import (
	PlaneService "Plane-Booking/app/service/plane"
	models "Plane-Booking/db/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddPlane(context *gin.Context) {

	var input models.Plane

	bindingErr := context.ShouldBindJSON(&input)

	if bindingErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": true,
			"error":   "Expected a plane type object in req body",
		})

		return
	}

	planeObject, err := PlaneService.AddPlane(&input)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": true,
			"error":   "Error in saving doc in DB",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"success": true,
		"user":    planeObject,
	})
}

func GetPlane(context *gin.Context) {

	var input = context.Query("id")

	plane, err := PlaneService.GetPlane(input)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": true,
			"error":   "Error in saving doc in DB",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"success": true,
		"user":    plane,
	})
}
