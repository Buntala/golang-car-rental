package controller

import (
	"car-rental/request"
	"car-rental/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MembershipController interface {
	FindAll(ctx *gin.Context) 
	FindOne(ctx *gin.Context)
	Save(ctx *gin.Context) 
	Update(ctx *gin.Context) 
	Delete(ctx *gin.Context)
}

type membershipController struct {
	service service.MembershipService
}

func NewMembership(service service.MembershipService) MembershipController {
	return &membershipController{
		service: service,
	}
}

func (c *membershipController) FindAll(ctx *gin.Context){
	res := c.service.FindAll()
	ctx.JSON(200, res)
}
func (c *membershipController) FindOne(ctx *gin.Context) { // entity.Membership
	body := request.Membership{}
	intID, err := strconv.Atoi(ctx.Param("id"))
	if err!=nil{
		panic(request.NewRequestError(err))
	}
	body.MembershipID = intID
	res, err := c.service.FindOne(body)
	if err!=nil{
		panic(request.NewRequestError(err))
	}
	ctx.JSON(200, res)
}

func (c *membershipController) Save(ctx *gin.Context) { // entity.Membership
	body := request.Membership{}
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

func (c *membershipController) Update(ctx *gin.Context) { // entity.Membership
	body := request.Membership{}
	err := ctx.ShouldBindJSON(&body)
	if err!=nil{
		panic(request.NewRequestError(err))
	}
	intID, err := strconv.Atoi(ctx.Param("id"))
	if err!=nil{
		panic(request.NewRequestError(err))
	}
	body.MembershipID = intID
	res, err := c.service.Update(body)
	if err!=nil{
		panic(request.NewRequestError(err))
	}
	ctx.JSON(200, res)
}

func (c *membershipController) Delete(ctx *gin.Context) { // entity.Membership
	body := request.Membership{}
	intID, err := strconv.Atoi(ctx.Param("id"))
	if err!=nil{
		panic(request.NewRequestError(err))
	}
	body.MembershipID = intID
	res, err := c.service.Delete(body)
	if err!=nil{
		panic(request.NewRequestError(err))
	}
	ctx.JSON(200, res)
}
