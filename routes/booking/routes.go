package booking

import "github.com/gin-gonic/gin"

func Booking(router *gin.Engine) {
	BookingRouter := router.Group("/bookings")
	BookingRouter.GET("/", getBooking)
	BookingRouter.GET("/:id", getBookingById)
	BookingRouter.POST("/", postBooking)
	BookingRouter.PATCH("/:id", patchBooking)
	BookingRouter.DELETE("/:id", deleteBooking)
	BookingRouter.POST("/:id/extend", extendBooking)
	//BookingRouter.POST("/:id/cancel", cancelBooking)
	BookingRouter.POST("/:id/finish", finishBooking)
}