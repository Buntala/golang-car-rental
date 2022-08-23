package driver

import (
	"car-rental/utilities/responseHandler"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)



func getDriver(c *gin.Context) {
	result := DBGetDriverAll()
	c.JSON(http.StatusOK, gin.H{
		"data" : result,
	})
}
func getDriverById(c *gin.Context) {
	body := DriverVal{}
	intVar, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responseHandler.ErrorHandler(errors.New("input id is not an integer"),c)
		return
	}
	body.DriverId = intVar
	body.Validate("get")
	result,err := DBGetDriverOne(body)
	if err!=nil{
		responseHandler.ErrorHandler(err,c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data" : result,
	})
}
func postDriver(c *gin.Context) {
	//conn := db.DbConnect()
	body := DriverVal{}
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
	if err := DBInsertDriver(&body); err!=nil{
		responseHandler.ErrorHandler(err,c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data" : body,
	})
}
func patchDriver(c *gin.Context) {
	//conn := db.DbConnect()
	body := DriverVal{}
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
	body.DriverId = intVar
	err = body.Validate("update")
	if err!=nil{
		responseHandler.ErrorHandler(err,c)
		return
	}
	if err := DBUpdateDriver(&body); err!=nil{
		responseHandler.ErrorHandler(err,c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data" : body,
	})
}
func deleteDriver(c *gin.Context) {
	//conn := db.DbConnect()
	body := DriverVal{}
	intVar, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responseHandler.ErrorHandler(errors.New("input id is not an integer"),c)
	}
	body.DriverId = intVar
	err = body.Validate("delete")
	if err!=nil{
		responseHandler.ErrorHandler(err,c)
		return
	}
	if err := DBDeleteDriver(&body); err!=nil{
		responseHandler.ErrorHandler(err,c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data" : body,
	})
}