package routes

import (
	controller "Plane-Booking/app/controllers/booking"

	"github.com/gin-gonic/gin"
)

func BookingRoutes(router *gin.Engine) {
	userRoutes := router.Group("/booking")

	userRoutes.POST("/", controller.Book)
	userRoutes.GET("/", controller.ShowBookings)
}
