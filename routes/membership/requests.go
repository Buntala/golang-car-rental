package membership

import (
	"log"
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
		panic("Input id is not an integer")
	}
	body.MembershipId = intVar
	body.Validate("get")
	result := DBGetMembershipOne(body)
	c.JSON(http.StatusOK, gin.H{
		"data" : result,
	})
}
func postMembership(c *gin.Context) {
	body := MembershipVal{}
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
	DBInsertMembership(&body)
	c.JSON(http.StatusOK, gin.H{
		"data" : body,
	})
}
func patchMembership(c *gin.Context) {
	//conn := db.DbConnect()
	body := MembershipVal{}
	intVar, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal("Input id is not an integer")
	}
	err = c.ShouldBindJSON(&body)
	if err!=nil{
		c.AbortWithError(http.StatusBadRequest,err)
		return
	}
	body.MembershipId = intVar
	if err != nil {
		log.Fatal(err)
	}
	DBUpdateMembership(&body)
	c.JSON(http.StatusOK, gin.H{
		"data" : body,
	})
}
func deleteMembership(c *gin.Context) {
	//conn := db.DbConnect()
	body := MembershipVal{}
	intVar, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal("Input id is not an integer")
	}
	body.MembershipId = intVar
	err = body.Validate("delete")
	if err!=nil{
		c.AbortWithError(http.StatusBadRequest,err)
		return
	}
	DBDeleteMembership(&body)
	c.JSON(http.StatusOK, gin.H{
		"data" : body,
	})
}