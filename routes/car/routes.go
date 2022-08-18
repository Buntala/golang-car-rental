package car

import "github.com/gin-gonic/gin"

func Car(router *gin.Engine) {
	carRouter := router.Group("/cars")
	carRouter.GET("/", getCar)
	carRouter.GET("/:id", getCarById)
	carRouter.POST("/", postCar)
	carRouter.PATCH("/:id", patchCar)
	carRouter.DELETE("/:id", deleteCar)
}