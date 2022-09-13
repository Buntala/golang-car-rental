package router

import (
	"car-rental/controller"
	"car-rental/repository"
	"car-rental/service"

	"github.com/gin-gonic/gin"
)

var (
	driverRepository repository.DriverRepository = repository.NewDriverRepository()
	driverService    service.DriverService       = service.NewDriverService(driverRepository)
	driverController controller.DriverController = controller.NewDriver(driverService)
)
func Driver(router *gin.Engine) {
	driver_r := router.Group("/drivers")
	{
		driver_r.GET("/", driverController.FindAll)
		driver_r.GET("/:id", driverController.FindOne)
		driver_r.POST("/",  driverController.Save)
		driver_r.PATCH("/:id",  driverController.Update)
		driver_r.DELETE("/:id",  driverController.Delete)
	}
}