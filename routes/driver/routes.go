package driver

import "github.com/gin-gonic/gin"

func Driver(router *gin.Engine) {
	router.GET("/drivers", getDriver)
	router.GET("/drivers/:id", getDriverById)
	router.POST("/drivers/", postDriver)
	router.PATCH("/drivers/:id", patchDriver)
	router.DELETE("/drivers/:id", deleteDriver)
}