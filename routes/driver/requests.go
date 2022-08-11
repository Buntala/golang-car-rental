package driver

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func getDriver(c *gin.Context) {
	//conn := db.DbConnect()
	result := DBGetDriverAll()
	c.JSON(http.StatusOK, gin.H{
		"data" : result,
	})
}
func getDriverById(c *gin.Context) {
	//conn := db.DbConnect()
	id := c.Param("id")
	result := DBGetDriverOne(id)
	c.JSON(http.StatusOK, gin.H{
		"data" : result,
	})
}
func postDriver(c *gin.Context) {
	//conn := db.DbConnect()
	body := Post_Rules{}
	if err:=c.ShouldBindBodyWith(&body,binding.JSON);err!=nil{
		c.AbortWithError(http.StatusBadRequest,err)
		return
	}
	validate := validator.New()
	err := validate.Struct(body)
	if err != nil {
		log.Fatal(err)
	}
	result := DBInsertDriver(body)
	c.JSON(http.StatusOK, gin.H{
		"data" : result,
	})
}
func patchDriver(c *gin.Context) {
	//conn := db.DbConnect()
	body := Patch_Rules{}
	if err:=c.ShouldBindBodyWith(&body,binding.JSON);err!=nil{
		c.AbortWithError(http.StatusBadRequest,err)
		return
	}
	body.DriverId = c.Param("id")
	validate := validator.New()
	err := validate.Struct(body)
	if err != nil {
		log.Fatal(err)
	}
	result := DBUpdateDriver(body)
	c.JSON(http.StatusOK, gin.H{
		"data" : result,
	})
}
func deleteDriver(c *gin.Context) {
	//conn := db.DbConnect()
	body := Delete_Rules{}
	body.DriverId = c.Param("id")
	validate := validator.New()
	err := validate.Struct(body)
	if err != nil {
		log.Fatal(err)
	}
	result := DBDeleteDriver(body)
	c.JSON(http.StatusOK, gin.H{
		"data" : result,
	})
}