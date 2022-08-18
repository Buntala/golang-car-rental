package driver

import "github.com/gin-gonic/gin"

func Driver(router *gin.Engine) {
	driverRouter := router.Group("/drivers")
	driverRouter.GET("/", getDriver)
	driverRouter.GET("/:id", getDriverById)
	driverRouter.POST("/", postDriver)
	driverRouter.PATCH("/:id", patchDriver)
	driverRouter.DELETE("/:id", deleteDriver)
}