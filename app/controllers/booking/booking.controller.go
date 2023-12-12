package controller

import (
	BookingService "Plane-Booking/app/service/book"
	models "Plane-Booking/db/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Book(context *gin.Context) {

	var input models.Booking

	bindingErr := context.ShouldBindJSON(&input)

	if bindingErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": true,
			"error":   "Expected a booking type object in req body",
		})

		return
	}

	booking, err := BookingService.BookSeat(&input)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": true,
			"error":   "Error in saving doc in DB",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"success": true,
		"user":    booking,
	})

}

func ShowBookings(context *gin.Context) {

	var input = context.Query("id")

	user, _ := BookingService.FindBooking(input)

	context.JSON(http.StatusCreated, gin.H{
		"success": true,
		"user":    user,
	})
}
