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
	customerController controller.CustomerController = controller.New(customerService)
)

func main() {
	server := gin.Default()

	server.GET("/customers", func(ctx *gin.Context) {
		ctx.JSON(200, customerController.FindAll())
	})
/*
	server.POST("/Customers", func(ctx *gin.Context) {
		ctx.JSON(200, CustomerController.Save(ctx))
	})
*/
	server.Run("127.0.0.1:8080")
}

