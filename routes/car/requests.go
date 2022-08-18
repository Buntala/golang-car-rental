package car

import (
	"log"
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
		panic("Input id is not an integer")
	}
	body.CarsID = intVar
	body.Validate("get")
	result := DBGetCarOne(body)
	c.JSON(http.StatusOK, gin.H{
		"data" : result,
	})
}
func postCar(c *gin.Context) {
	//conn := db.DbConnect()
	body := CarDB{}
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
	DBInsertCar(&body)
	c.JSON(http.StatusOK, gin.H{
		"data" : body,
	})
}
func patchCar(c *gin.Context) {
	//conn := db.DbConnect()
	body := CarDB{}
	intVar, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal("Input id is not an integer")
	}
	err = c.ShouldBindJSON(&body)
	if err!=nil{
		c.AbortWithError(http.StatusBadRequest,err)
		return
	}
	body.CarsID = intVar
	err = body.Validate("update")
	if err!=nil{
		c.AbortWithError(http.StatusBadRequest,err)
		return
	}
	DBUpdateCar(&body)
	c.JSON(http.StatusOK, gin.H{
		"data" : body,
	})
}
func deleteCar(c *gin.Context) {
	//conn := db.DbConnect()
	body := CarDB{}
	intVar, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal("Input id is not an integer")
	}
	body.CarsID = intVar
	err = body.Validate("delete")
	if err!=nil{
		c.AbortWithError(http.StatusBadRequest,err)
		return
	}
	DBDeleteCar(&body)
	c.JSON(http.StatusOK, gin.H{
		"data" : body,
	})
}