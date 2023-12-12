package routes

import "github.com/gin-gonic/gin"

func Routes() *gin.Engine {
	router := gin.Default()

	UserRoutes(router)
	BookingRoutes(router)
	PlaneRoutes(router)

	return router

}
