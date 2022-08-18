package customer

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getCustomer(c *gin.Context) {
	result := DBGetCustomerAll()
	c.JSON(http.StatusOK, gin.H{
		"data" : result,
	})
}
func getCustomerById(c *gin.Context) {
	body := CustomerDB{}
	intVar, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic("Input id is not an integer")
	}
	body.CustomerID = intVar
	body.Validate("get")
	result := DBGetCustomerOne(body)
	c.JSON(http.StatusOK, gin.H{
		"data" : result,
	})
}
func postCustomer(c *gin.Context) {
	//conn := db.DbConnect()
	body := CustomerDB{}
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
	DBInsertCustomer(&body)
	c.JSON(http.StatusOK, gin.H{
		"data" : body,
	})
}
func patchCustomer(c *gin.Context) {
	//conn := db.DbConnect()
	body := CustomerDB{}
	intVar, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal("Input id is not an integer")
	}
	err = c.ShouldBindJSON(&body)
	if err!=nil{
		c.AbortWithError(http.StatusBadRequest,err)
		return
	}
	body.CustomerID = intVar
	err = body.Validate("update")
	if err!=nil{
		c.AbortWithError(http.StatusBadRequest,err)
		return
	}
	DBUpdateCustomer(&body)
	c.JSON(http.StatusOK, gin.H{
		"data" : body,
	})
}
func deleteCustomer(c *gin.Context) {
	//conn := db.DbConnect()
	body := CustomerDB{}
	intVar, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal("Input id is not an integer")
	}
	body.CustomerID = intVar
	err = body.Validate("delete")
	if err!=nil{
		c.AbortWithError(http.StatusBadRequest,err)
		return
	}
	DBDeleteCustomer(&body)
	c.JSON(http.StatusOK, gin.H{
		"data" : body,
	})
}