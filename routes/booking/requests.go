package booking

import (
	"car-rental/utilities/responseHandler"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)



func getBooking(c *gin.Context) {
	result := DBGetBookingAll()
	c.JSON(http.StatusOK, gin.H{
		"data" : result,
	})
}
func getBookingById(c *gin.Context) {
	body := BookingDB{}
	intVar, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responseHandler.ErrorHandler(errors.New("input id is not an integer"),c)
		return
	}
	body.BookingID = intVar
	body.Validate("get")
	result,err := DBGetBookingOne(&body)
	if err!=nil{
		responseHandler.ErrorHandler(err,c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data" : result,
	})
}
func postBooking(c *gin.Context) {
	body := BookingDB{}
	err := c.ShouldBindJSON(&body)
	if err!=nil{
		c.AbortWithError(http.StatusBadRequest,err)
		return
	}
	err = body.Validate("post")
	if err!=nil{
		c.AbortWithError(http.StatusBadRequest,err)
		return
	}
	//result := 
	if err := DBInsertBooking(&body); err!=nil{
		responseHandler.ErrorHandler(err,c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data" : body,
	})
}
func patchBooking(c *gin.Context) {
	body := BookingDB{}
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
	body.BookingID = intVar
	err = body.Validate("update")
	if err!=nil{
		responseHandler.ErrorHandler(err,c)
		return
	}
	if err := DBUpdateBooking(&body); err!=nil{
		responseHandler.ErrorHandler(err,c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data" : body,
	})
}
func deleteBooking(c *gin.Context) {
	body := BookingDB{}
	intVar, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responseHandler.ErrorHandler(errors.New("input id is not an integer"),c)
		return 
	}
	body.BookingID = intVar
	err = body.Validate("delete")
	if err!=nil{
		responseHandler.ErrorHandler(err,c)
		return
	}
	if err := DBDeleteBooking(&body); err!=nil{
		responseHandler.ErrorHandler(err,c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data" : body,
	})
}
func extendBooking(c *gin.Context) {
	body := BookingDB{}
	intVar, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responseHandler.ErrorHandler(errors.New("input id is not an integer"),c)
		return
	}
	body.BookingID = intVar
	err = c.ShouldBindJSON(&body)
	if err!=nil{
		responseHandler.ErrorHandler(err,c)
		return
	}
	err = body.Validate("extend")
	if err!=nil{
		c.AbortWithError(http.StatusBadRequest,err)
		return
	}
	if err := DBExtend(&body); err!=nil{
		responseHandler.ErrorHandler(err,c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data" : body,
	})
}
/*
func cancelBooking(c *gin.Context) {
	//conn := db.DbConnect()
	body := BookingDB{}
	intVar, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responseHandler.ErrorHandler(errors.New("input id is not an integer"),c)
	}
	body.BookingID = intVar
	err = body.Validate("cancel")
	if err!=nil{
		c.AbortWithError(http.StatusBadRequest,err)
		return
	}
	if err := DBCanceled(&body); err!=nil{
		responseHandler.ErrorHandler(err,c)
	}
	c.JSON(http.StatusOK, gin.H{
		"data" : body,
	})
}*/
func finishBooking(c *gin.Context) {
	//conn := db.DbConnect()
	body := BookingDB{}
	intVar, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responseHandler.ErrorHandler(errors.New("input id is not an integer"),c)
	}
	body.BookingID = intVar
	err = body.Validate("finish")
	if err!=nil{
		c.AbortWithError(http.StatusBadRequest,err)
		return
	}
	if err := DBFinished(&body); err!=nil{
		responseHandler.ErrorHandler(err,c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data" : body,
	})
}