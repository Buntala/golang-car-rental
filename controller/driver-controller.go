package controller

import (
	"car-rental/request"
	"car-rental/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DriverController interface {
	FindAll(ctx *gin.Context) 
	FindOne(ctx *gin.Context) 
	Save(ctx *gin.Context) 
	Update(ctx *gin.Context) 
	Delete(ctx *gin.Context)
}

type driverController struct {
	service service.DriverService
}

func NewDriver(service service.DriverService) DriverController {
	return &driverController{
		service: service,
	}
}

func (c *driverController) FindAll(ctx *gin.Context){
	ctx.JSON(200, c.service.FindAll())
}

func (c *driverController) FindOne(ctx *gin.Context) {
	body := request.Driver{}
	intID, err := strconv.Atoi(ctx.Param("id"))
	if err!=nil{
		panic(request.NewRequestError(err))
	}
	body.DriverID = intID
	res,err:= c.service.FindOne(body)
	if err!=nil{
		panic(request.NewRequestError(err))
	}
	ctx.JSON(200, res)
}

func (c *driverController) Save(ctx *gin.Context) {
	body := request.Driver{}
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

func (c *driverController) Update(ctx *gin.Context) {
	body := request.Driver{}
	err := ctx.ShouldBindJSON(&body)
	if err!=nil{
		panic(request.NewRequestError(err))
	}
	intID, err := strconv.Atoi(ctx.Param("id"))
	body.DriverID = intID
	
	if err!=nil{
		panic(request.NewRequestError(err))
	}
	res ,err := c.service.Update(body)
	if err!=nil{
		panic(request.NewRequestError(err))
	}
	ctx.JSON(200, res)
}

func (c *driverController) Delete(ctx *gin.Context){
	body := request.Driver{}
	err := ctx.ShouldBindJSON(&body)
	if err!=nil{
		panic(request.NewRequestError(err))
	}
	intID, err := strconv.Atoi(ctx.Param("id"))
	body.DriverID = intID
	
	if err!=nil{
		panic(request.NewRequestError(err))
	}
	res,err := c.service.Delete(body)
	if err!=nil{
		panic(request.NewRequestError(err))
	}
	ctx.JSON(200, res)
}
