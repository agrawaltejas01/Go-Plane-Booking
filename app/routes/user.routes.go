package routes

import (
	controller "Plane-Booking/app/controllers/user"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	userRoutes := router.Group("/user")

	userRoutes.POST("/", controller.AddUser)
	userRoutes.GET("/", controller.GetUser)
}
