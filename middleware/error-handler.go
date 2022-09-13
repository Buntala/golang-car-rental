package middleware

import (
	"car-rental/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorResponse(c *gin.Context, err any){
	c.JSON(http.StatusOK, gin.H{
		"error": err.(request.RequestError).StatusCode,
		"message" : err.(request.RequestError).Error(),
	})
}
