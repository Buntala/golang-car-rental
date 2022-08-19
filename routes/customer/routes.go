package customer

import "github.com/gin-gonic/gin"

func Customer(router *gin.Engine) {
	customerRouter := router.Group("/customers")
	customerRouter.GET("/", getCustomer)
	customerRouter.GET("/:id", getCustomerById)
	customerRouter.POST("/", postCustomer)
	customerRouter.PATCH("/:id", patchCustomer)
	customerRouter.DELETE("/:id", deleteCustomer)
}