package membership

import (
	"car-rental/utilities/responseHandler"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getMembership(c *gin.Context) {
	result := DBGetMembershipAll()
	c.JSON(http.StatusOK, gin.H{
		"data" : result,
	})
}

///not changed
func getMembershipById(c *gin.Context) {
	body := MembershipVal{}
	intVar, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responseHandler.ErrorHandler(errors.New("input id is not an integer"),c)
	}
	body.MembershipId = intVar
	if err := body.Validate("get"); err!=nil{
		responseHandler.ErrorHandler(err,c)
	}
	result,err := DBGetMembershipOne(body)
	if err!=nil{
		responseHandler.ErrorHandler(err,c)
	}
	c.JSON(http.StatusOK, gin.H{
		"data" : result,
	})
}
func postMembership(c *gin.Context) {
	body := MembershipVal{}
	err := c.ShouldBindJSON(&body)
	if err!=nil{
		responseHandler.ErrorHandler(err,c)
	}
	err = body.Validate("post")
	if err!=nil{
		responseHandler.ErrorHandler(err,c)
	}
	if err := DBInsertMembership(&body); err!=nil{
		responseHandler.ErrorHandler(err,c)
	}
	c.JSON(http.StatusOK, gin.H{
		"data" : body,
	})
}
func patchMembership(c *gin.Context) {
	//conn := db.DbConnect()
	body := MembershipVal{}
	intVar, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responseHandler.ErrorHandler(errors.New("input id is not an integer"),c)
	}
	err = c.ShouldBindJSON(&body)
	if err!=nil{
		responseHandler.ErrorHandler(err,c)
	}
	body.MembershipId = intVar
	if err!=nil{
		responseHandler.ErrorHandler(err,c)
	}
	if err := DBUpdateMembership(&body); err!=nil{
		responseHandler.ErrorHandler(err,c)
	}
	c.JSON(http.StatusOK, gin.H{
		"data" : body,
	})
}
func deleteMembership(c *gin.Context) {
	//conn := db.DbConnect()
	body := MembershipVal{}
	intVar, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responseHandler.ErrorHandler(errors.New("input id is not an integer"),c)
	}
	body.MembershipId = intVar
	err = body.Validate("delete")
	if err!=nil{
		responseHandler.ErrorHandler(err,c)
	}
	if err := DBDeleteMembership(&body); err!=nil{
		responseHandler.ErrorHandler(err,c)
	}
	c.JSON(http.StatusOK, gin.H{
		"data" : body,
	})
}