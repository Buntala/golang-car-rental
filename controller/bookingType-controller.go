package controller

import (
	"car-rental/request"
	"car-rental/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookingTypeController interface {
	FindAll(ctx *gin.Context) 
	FindOne(ctx *gin.Context)
	Save(ctx *gin.Context) 
	Update(ctx *gin.Context) 
	Delete(ctx *gin.Context)
}

type bookingTypeController struct {
	service service.BookingTypeService
}

func NewBookingType(service service.BookingTypeService) BookingTypeController {
	return &bookingTypeController{
		service: service,
	}
}

func (c *bookingTypeController) FindAll(ctx *gin.Context){
	res := c.service.FindAll()
	ctx.JSON(200, res)
}
func (c *bookingTypeController) FindOne(ctx *gin.Context) { // entity.BookingType
	body := request.BookingType{}
	intID, err := strconv.Atoi(ctx.Param("id"))
	if err!=nil{
		panic(request.NewRequestError(err))
	}
	body.BookingTypeID = intID
	res, err := c.service.FindOne(body)
	if err!=nil{
		panic(request.NewRequestError(err))
	}
	ctx.JSON(200, res)
}

func (c *bookingTypeController) Save(ctx *gin.Context) { // entity.BookingType
	body := request.BookingType{}
	err := ctx.ShouldBindJSON(&body)
	if err!=nil{
		panic(request.NewRequestError(err))
	}
	res, err := c.service.Save(body)
	if err!=nil{
		panic(request.NewRequestError(err))
	}
	ctx.JSON(200, res)
}

func (c *bookingTypeController) Update(ctx *gin.Context) { // entity.BookingType
	body := request.BookingType{}
	err := ctx.ShouldBindJSON(&body)
	if err!=nil{
		panic(request.NewRequestError(err))
	}
	intID, err := strconv.Atoi(ctx.Param("id"))
	if err!=nil{
		panic(request.NewRequestError(err))
	}
	body.BookingTypeID = intID
	res, err := c.service.Update(body)
	if err!=nil{
		panic(request.NewRequestError(err))
	}
	ctx.JSON(200, res)
}

func (c *bookingTypeController) Delete(ctx *gin.Context) { // entity.BookingType
	body := request.BookingType{}
	intID, err := strconv.Atoi(ctx.Param("id"))
	if err!=nil{
		panic(request.NewRequestError(err))
	}
	body.BookingTypeID = intID
	res, err := c.service.Delete(body)
	if err!=nil{
		panic(request.NewRequestError(err))
	}
	ctx.JSON(200, res)
}
