package main

import (
	"car-rental/controller"
	"car-rental/repository"
	"car-rental/service"

	"github.com/gin-gonic/gin"
)

var (
	customerRepository repository.CustomerRepository = repository.NewCustomerRepository()
	customerService    service.CustomerService = service.NewCustomerService(customerRepository)
	customerController controller.CustomerController = controller.NewCustomer(customerService)

	membershipRepository repository.MembershipRepository = repository.NewMembershipRepository()
	membershipService    service.MembershipService = service.NewMembershipService(membershipRepository)
	membershipController controller.MembershipController = controller.NewMembership(membershipService)
	
	carRepository repository.CarRepository = repository.NewCarRepository()
	carService    service.CarService = service.NewCarService(carRepository)
	carController controller.CarController = controller.NewCar(carService)
	
)

func main() {
	server := gin.Default()
	
	cust_r := server.Group("/customers")
	cust_r.GET("/", customerController.FindAll)
	cust_r.GET("/:id", customerController.FindOne)
	cust_r.POST("/", customerController.Save)
	cust_r.PATCH("/:id",  customerController.Update)
	cust_r.DELETE("/:id", customerController.Delete)
	cust_r.POST("/:id/membership", customerController.SaveMembership)

	
	member_r := server.Group("/memberships")

	member_r.GET("/", membershipController.FindAll)
	member_r.GET("/:id", membershipController.FindOne)
	member_r.POST("/",  membershipController.Save)
	member_r.PATCH("/:id",  membershipController.Update)
	member_r.DELETE("/:id",  membershipController.Delete)

	car_r := server.Group("/cars")

	car_r.GET("/", carController.FindAll)
	car_r.GET("/:id", carController.FindOne)
	car_r.POST("/",  carController.Save)
	car_r.PATCH("/:id",  carController.Update)
	car_r.DELETE("/:id",  carController.Delete)

	server.Run("127.0.0.1:8080")
}

