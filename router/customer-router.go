package router

import (
	"car-rental/controller"
	"car-rental/repository"
	"car-rental/service"

	"github.com/gin-gonic/gin"
)

var (
	customerRepository repository.CustomerRepository = repository.NewCustomerRepository()
	customerService    service.CustomerService       = service.NewCustomerService(customerRepository)
	customerController controller.CustomerController = controller.NewCustomer(customerService)
)

func Customer(router *gin.Engine) {
	cust_r := router.Group("/customers")
	{
		cust_r.GET("/", customerController.FindAll)
		cust_r.GET("/:id", customerController.FindOne)
		cust_r.POST("/", customerController.Save)
		cust_r.PATCH("/:id", customerController.Update)
		cust_r.DELETE("/:id", customerController.Delete)
		cust_r.POST("/:id/membership", customerController.SaveMembership)

	}
}