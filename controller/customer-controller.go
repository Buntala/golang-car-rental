package controller

import (
	"car-rental/request"
	"car-rental/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CustomerController interface {
	FindAll(ctx *gin.Context) 
	FindOne(ctx *gin.Context) 
	Save(ctx *gin.Context) 
	Update(ctx *gin.Context) 
	Delete(ctx *gin.Context)
	SaveMembership(ctx *gin.Context)
}

type customerController struct {
	service service.CustomerService
}

func NewCustomer(service service.CustomerService) CustomerController {
	return &customerController{
		service: service,
	}
}

func (c *customerController) FindAll(ctx *gin.Context){
	ctx.JSON(200, c.service.FindAll())
}

func (c *customerController) FindOne(ctx *gin.Context) {
	body := request.CustomerRequest{}
	intID, err := strconv.Atoi(ctx.Param("id"))
	if err!=nil{
		panic(request.NewRequestError(err))
	}
	body.CustomerID = intID
	res,err:= c.service.FindOne(body)
	if err!=nil{
		panic(request.NewRequestError(err))
	}
	ctx.JSON(200, res)
}

func (c *customerController) Save(ctx *gin.Context) {
	body := request.CustomerRequest{}
	err := ctx.ShouldBindJSON(&body)
	if err!=nil{
		panic(request.NewRequestError(err))
	}
	res,err :=c.service.Save(body)
	if err!=nil{
		panic(request.NewRequestError(err))
	}
	ctx.JSON(200, res)
}

func (c *customerController) Update(ctx *gin.Context) {
	body := request.CustomerRequest{}
	err := ctx.ShouldBindJSON(&body)
	if err!=nil{
		panic(request.NewRequestError(err))
	}
	intID, err := strconv.Atoi(ctx.Param("id"))
	body.CustomerID = intID
	
	if err!=nil{
		panic(request.NewRequestError(err))
	}
	res ,err := c.service.Update(body)
	if err!=nil{
		panic(request.NewRequestError(err))
	}
	ctx.JSON(200, res)
}

func (c *customerController) Delete(ctx *gin.Context){
	body := request.CustomerRequest{}
	err := ctx.ShouldBindJSON(&body)
	if err!=nil{
		panic(request.NewRequestError(err))
	}
	intID, err := strconv.Atoi(ctx.Param("id"))
	body.CustomerID = intID
	
	if err!=nil{
		panic(request.NewRequestError(err))
	}
	res,err := c.service.Delete(body)
	if err!=nil{
		panic(request.NewRequestError(err))
	}
	ctx.JSON(200, res)
}


func (c *customerController) SaveMembership(ctx *gin.Context) {
	body := request.CustomerRequest{}
	err := ctx.ShouldBindJSON(&body)
	if err!=nil{
		panic(request.NewRequestError(err))
	}
	intID, err := strconv.Atoi(ctx.Param("id"))
	body.CustomerID = intID
	if err!=nil{
		panic(request.NewRequestError(err))
	}
	res,err:= c.service.SaveMembership(body)
	if err!=nil{
		panic(request.NewRequestError(err))
	}
	
	ctx.JSON(200, res)
}
