package customer

import (
	"car-rental/utilities/responseHandler"
	"errors"
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
		responseHandler.ErrorHandler(errors.New("input id is not an integer"),c)
		return
	}
	body.CustomerID = intVar
	body.Validate("get")
	result, err := DBGetCustomerOne(body)
	if err != nil {		
		responseHandler.ErrorHandler(err,c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data" : result,
	})
}
func postCustomer(c *gin.Context) {
	//conn := db.DbConnect()
	body := CustomerDB{}
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
	DBInsertCustomer(&body)
	c.JSON(http.StatusOK, gin.H{
		"data" : body,
	})
}
func patchCustomer(c *gin.Context) {
	//conn := db.DbConnect()
	body := CustomerDB{}
	intVar, err := strconv.Atoi(c.Param("id"))
	if err!=nil{
		responseHandler.ErrorHandler(errors.New("input id is not an integer"),c)
		return
	}
	err = c.ShouldBindJSON(&body)
	if err!=nil{
		responseHandler.ErrorHandler(err,c)
		return
	}
	body.CustomerID = intVar
	err = body.Validate("update")
	if err!=nil{
		responseHandler.ErrorHandler(err,c)
		return
	}
	if err:=DBUpdateCustomer(&body);err!=nil{
		responseHandler.ErrorHandler(err,c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data" : body,
	})
}
func deleteCustomer(c *gin.Context) {
	//conn := db.DbConnect()
	body := CustomerDB{}
	intVar, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responseHandler.ErrorHandler(errors.New("input id is not an integer"),c)
		return
	}
	body.CustomerID = intVar
	err = body.Validate("delete")
	if err!=nil{
		responseHandler.ErrorHandler(err,c)
		return
	}
	if err := DBDeleteCustomer(&body); err!=nil{
		responseHandler.ErrorHandler(err,c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data" : body,
	})
}
func postMembership(c *gin.Context) {
	//conn := db.DbConnect()
	body := CustomerDB{}
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
	body.CustomerID = intVar
	if err := DBUpdateMembershipStatus(&body); err!=nil{
		responseHandler.ErrorHandler(err,c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data" : body,
	})
}