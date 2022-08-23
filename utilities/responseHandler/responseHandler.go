package responseHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(e error,c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error":"400 Bad Request",
		"message" : e.Error(),
	})
}
