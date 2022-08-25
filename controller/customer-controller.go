package controller

import (
	"car-rental/entity"
	"car-rental/service"

	"github.com/gin-gonic/gin"
)

type CustomerController interface {
	FindAll() []entity.Customer
	Save(ctx *gin.Context) entity.Customer
}

type controller struct {
	service service.CustomerService
}

func New(service service.CustomerService) CustomerController {
	return &controller{
		service: service,
	}
}


func (c *controller) FindAll() []entity.Customer {
	res := c.service.FindAll()
	return res
}
func (c *controller) Save(ctx *gin.Context) entity.Customer {
	body := entity.Customer{}
	err := ctx.ShouldBindJSON(&body)
	if err!=nil{
		panic(err)
	}

	return body//CHANGE LATER
}