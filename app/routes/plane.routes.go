package routes

import (
	controller "Plane-Booking/app/controllers/plane"

	"github.com/gin-gonic/gin"
)

func PlaneRoutes(router *gin.Engine) {
	userRoutes := router.Group("/plane")

	userRoutes.POST("/", controller.AddPlane)
	userRoutes.GET("/", controller.GetPlane)
}
