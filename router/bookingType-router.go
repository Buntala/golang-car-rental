package router

import (
	"car-rental/controller"
	"car-rental/repository"
	"car-rental/service"

	"github.com/gin-gonic/gin"
)

var (
	bookingTypeRepository repository.BookingTypeRepository = repository.NewBookingTypeRepository()
	bookingTypeService    service.BookingTypeService       = service.NewBookingTypeService(bookingTypeRepository)
	bookingTypeController controller.BookingTypeController = controller.NewBookingType(bookingTypeService)
)

func BookingType(router *gin.Engine) {
	bookingtype_r := router.Group("/booking-type")
	{
		bookingtype_r.GET("/", bookingTypeController.FindAll)
		bookingtype_r.GET("/:id", bookingTypeController.FindOne)
		bookingtype_r.POST("/", bookingTypeController.Save)
		bookingtype_r.PATCH("/:id", bookingTypeController.Update)
		bookingtype_r.DELETE("/:id", bookingTypeController.Delete)
	}
}