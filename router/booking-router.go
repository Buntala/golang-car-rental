package router

import (
	"car-rental/controller"
	"car-rental/repository"
	"car-rental/service"

	"github.com/gin-gonic/gin"
)

var (
	bookingRepository repository.BookingRepository = repository.NewBookingRepository()
	bookingService    service.BookingService       = service.NewBookingService(bookingRepository)
	bookingController controller.BookingController = controller.NewBooking(bookingService)
)

func Booking(router *gin.Engine) {
	booking_r := router.Group("/bookings")
	{
		booking_r.GET("/", bookingController.FindAll)
		booking_r.GET("/:id", bookingController.FindOne)
		booking_r.POST("/", bookingController.Save)
		booking_r.PATCH("/:id", bookingController.Update)
		booking_r.DELETE("/:id", bookingController.Delete)
		booking_r.POST("/:id/finish", bookingController.SaveFinished)
		booking_r.POST("/:id/extend", bookingController.SaveExtend)
	}
}