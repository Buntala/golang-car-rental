package bookingType

import "github.com/gin-gonic/gin"

func BookingType(router *gin.Engine) {
	bookingTypeRouter := router.Group("/booking-types")
	bookingTypeRouter.GET("/", getBookingType)
	bookingTypeRouter.GET("/:id", getBookingTypeById)
	bookingTypeRouter.POST("/", postBookingType)
	bookingTypeRouter.PATCH("/:id", patchBookingType)
	bookingTypeRouter.DELETE("/:id", deleteBookingType)
}