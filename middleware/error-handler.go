package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorResponse(c *gin.Context, err any){
	c.JSON(http.StatusBadRequest, gin.H{
		"error":"400 Bad Request",
		"message" : err,
	})
}
