package car

import (
	"car-rental/utilities/responseHandler"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)



func getCar(c *gin.Context) {
	result := DBGetCarAll()
	c.JSON(http.StatusOK, gin.H{
		"data" : result,
	})
}
func getCarById(c *gin.Context) {
	body := CarDB{}
	intVar, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responseHandler.ErrorHandler(errors.New("input id is not an integer"),c)
		return
	}
	body.CarsID = intVar
	body.Validate("get")
	result,err := DBGetCarOne(body)
	if err!=nil{
		responseHandler.ErrorHandler(err,c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data" : result,
	})
}
func postCar(c *gin.Context) {
	//conn := db.DbConnect()
	body := CarDB{}
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
	if err := DBInsertCar(&body); err!=nil{
		responseHandler.ErrorHandler(err,c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data" : body,
	})
}
func patchCar(c *gin.Context) {
	//conn := db.DbConnect()
	body := CarDB{}
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
	body.CarsID = intVar
	err = body.Validate("update")
	if err!=nil{
		responseHandler.ErrorHandler(err,c)
		return
	}
	if err := DBUpdateCar(&body); err!=nil{
		responseHandler.ErrorHandler(err,c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data" : body,
	})
}
func deleteCar(c *gin.Context) {
	//conn := db.DbConnect()
	body := CarDB{}
	intVar, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responseHandler.ErrorHandler(errors.New("input id is not an integer"),c)
		return
	}
	body.CarsID = intVar
	err = body.Validate("delete")
	if err!=nil{
		responseHandler.ErrorHandler(err,c)
		return
	}
	if err := DBDeleteCar(&body); err!=nil{
		responseHandler.ErrorHandler(err,c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data" : body,
	})
}