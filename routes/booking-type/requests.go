package bookingType

import (
	"car-rental/utilities/responseHandler"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)



func getBookingType(c *gin.Context) {
	result := DBGetBookingTypeAll()
	c.JSON(http.StatusOK, gin.H{
		"data" : result,
	})
}
func getBookingTypeById(c *gin.Context) {
	body := BookingTypeDB{}
	intVar, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responseHandler.ErrorHandler(errors.New("input id is not an integer"),c)
		return
	}
	body.BookingTypeID = intVar
	body.Validate("get")
	result,err := DBGetBookingTypeOne(body)
	if err!=nil{
		responseHandler.ErrorHandler(err,c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data" : result,
	})
}
func postBookingType(c *gin.Context) {
	//conn := db.DbConnect()
	body := BookingTypeDB{}
	err := c.ShouldBindJSON(&body)
	if err!=nil{
		responseHandler.ErrorHandler(err,c)
		return
	}
	err = body.Validate("post")
	if err!=nil{
		responseHandler.ErrorHandler(err,c)
		return
	}
	//result := 
	if err := DBInsertBookingType(&body); err!=nil{
		responseHandler.ErrorHandler(err,c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data" : body,
	})
}
func patchBookingType(c *gin.Context) {
	//conn := db.DbConnect()
	body := BookingTypeDB{}
	intVar, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responseHandler.ErrorHandler(errors.New("input id is not an integer"),c)
		return
	}
	err = c.ShouldBindJSON(&body)
	if err!=nil{
		responseHandler.ErrorHandler(err,c)
		return
	}
	body.BookingTypeID = intVar
	err = body.Validate("update")
	if err!=nil{
		responseHandler.ErrorHandler(err,c)
		return
	}
	if err := DBUpdateBookingType(&body); err!=nil{
		responseHandler.ErrorHandler(err,c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data" : body,
	})
}
func deleteBookingType(c *gin.Context) {
	//conn := db.DbConnect()
	body := BookingTypeDB{}
	intVar, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responseHandler.ErrorHandler(errors.New("input id is not an integer"),c)
		return
	}
	body.BookingTypeID = intVar
	err = body.Validate("delete")
	if err!=nil{
		responseHandler.ErrorHandler(err,c)
		return
	}
	if err := DBDeleteBookingType(&body); err!=nil{
		responseHandler.ErrorHandler(err,c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data" : body,
	})
}