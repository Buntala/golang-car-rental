package router

import (
	"car-rental/controller"
	"car-rental/repository"
	"car-rental/service"

	"github.com/gin-gonic/gin"
)

var (
	carRepository repository.CarRepository = repository.NewCarRepository()
	carService    service.CarService       = service.NewCarService(carRepository)
	carController controller.CarController = controller.NewCar(carService)
)
func Car(router *gin.Engine) {
	car_r := router.Group("/cars")
	{
		car_r.GET("/", carController.FindAll)
		car_r.GET("/:id", carController.FindOne)
		car_r.POST("/",  carController.Save)
		car_r.PATCH("/:id",  carController.Update)
		car_r.DELETE("/:id",  carController.Delete)
	}
}