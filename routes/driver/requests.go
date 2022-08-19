package driver

import (
	"log"
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
		panic("Input id is not an integer")
	}
	body.DriverId = intVar
	body.Validate("get")
	result := DBGetDriverOne(body)
	c.JSON(http.StatusOK, gin.H{
		"data" : result,
	})
}
func postDriver(c *gin.Context) {
	//conn := db.DbConnect()
	body := DriverVal{}
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
	DBInsertDriver(&body)
	c.JSON(http.StatusOK, gin.H{
		"data" : body,
	})
}
func patchDriver(c *gin.Context) {
	//conn := db.DbConnect()
	body := DriverVal{}
	intVar, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal("Input id is not an integer")
	}
	err = c.ShouldBindJSON(&body)
	if err!=nil{
		c.AbortWithError(http.StatusBadRequest,err)
		return
	}
	body.DriverId = intVar
	err = body.Validate("update")
	if err!=nil{
		c.AbortWithError(http.StatusBadRequest,err)
		return
	}
	DBUpdateDriver(&body)
	c.JSON(http.StatusOK, gin.H{
		"data" : body,
	})
}
func deleteDriver(c *gin.Context) {
	//conn := db.DbConnect()
	body := DriverVal{}
	intVar, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal("Input id is not an integer")
	}
	body.DriverId = intVar
	err = body.Validate("delete")
	if err!=nil{
		c.AbortWithError(http.StatusBadRequest,err)
		return
	}
	DBDeleteDriver(&body)
	c.JSON(http.StatusOK, gin.H{
		"data" : body,
	})
}