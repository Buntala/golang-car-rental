package controller

import (
	"car-rental/request"
	"car-rental/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookingController interface {
	FindAll(ctx *gin.Context) 
	FindOne(ctx *gin.Context)
	Save(ctx *gin.Context) 
	Update(ctx *gin.Context) 
	Delete(ctx *gin.Context)
	SaveExtend(ctx *gin.Context)
	SaveFinished(ctx *gin.Context)
}

type bookingController struct {
	service service.BookingService
}

func NewBooking(service service.BookingService) BookingController {
	return &bookingController{
		service: service,
	}
}

func (c *bookingController) FindAll(ctx *gin.Context){
	res := c.service.FindAll()
	ctx.JSON(200, res)
}
func (c *bookingController) FindOne(ctx *gin.Context) { 
	body := request.Booking{}
	intID, err := strconv.Atoi(ctx.Param("id"))
	if err!=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":"400 Bad Request",
			"message" : err.Error(),
		})
		panic(err)
	}
	body.BookingID = intID
	res, err := c.service.FindOne(body)
	if err!=nil{
		/*
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":"400 Bad Request",
			"message" : err.Error(),
		})*/
		panic(err)
	}
	ctx.JSON(200, res)
}

func (c *bookingController) Save(ctx *gin.Context) { 
	body := request.Booking{}
	err := ctx.ShouldBindJSON(&body)
	if err!=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":"400 Bad Request",
			"message" : err.Error(),
		})
		return
	}
	res, err := c.service.Save(body)
	if err!=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":"400 Bad Request",
			"message" : err.Error(),
		})
		return
	}
	ctx.JSON(200, res)
}

func (c *bookingController) Update(ctx *gin.Context) { 
	body := request.Booking{}
	err := ctx.ShouldBindJSON(&body)
	if err!=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":"400 Bad Request",
			"message" : err.Error(),
		})
		return
	}
	intID, err := strconv.Atoi(ctx.Param("id"))
	if err!=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":"400 Bad Request",
			"message" : err.Error(),
		})
		return
	}
	body.BookingID = intID
	res, err := c.service.Update(body)
	if err!=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":"400 Bad Request",
			"message" : err.Error(),
		})
		return
	}
	ctx.JSON(200, res)
}

func (c *bookingController) Delete(ctx *gin.Context) { 
	body := request.Booking{}
	intID, err := strconv.Atoi(ctx.Param("id"))
	if err!=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":"400 Bad Request",
			"message" : err.Error(),
		})
		return
	}
	body.BookingID = intID
	res, err := c.service.Delete(body)
	if err!=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":"400 Bad Request",
			"message" : err.Error(),
		})
		return
	}
	ctx.JSON(200, res)
}

func (c *bookingController) SaveExtend(ctx *gin.Context) { 
	body := request.Booking{}
	err := ctx.ShouldBindJSON(&body)
	if err!=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":"400 Bad Request",
			"message" : err.Error(),
		})
		return
	}
	intID, err := strconv.Atoi(ctx.Param("id"))
	if err!=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":"400 Bad Request",
			"message" : err.Error(),
		})
		return
	}
	body.BookingID = intID
	res, err := c.service.SaveExtend(body)
	if err!=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":"400 Bad Request",
			"message" : err.Error(),
		})
		return
	}
	ctx.JSON(200, res)
}

func (c *bookingController) SaveFinished(ctx *gin.Context) { 
	body := request.Booking{}
	intID, err := strconv.Atoi(ctx.Param("id"))
	if err!=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":"400 Bad Request",
			"message" : err.Error(),
		})
		return
	}
	body.BookingID = intID
	res, err := c.service.SaveFinished(body)
	if err!=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":"400 Bad Request",
			"message" : err.Error(),
		})
		return
	}
	ctx.JSON(200, res)
}