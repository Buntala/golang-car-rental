package main

import (
	"car-rental/middleware"

	"github.com/gin-gonic/gin"
)
func main() {
	server := gin.New()
	server.Use(gin.Logger())
	server.Use(gin.CustomRecovery(middleware.ErrorResponse))
	//router := server.Group("/v1")
	Router(server)

	server.Run("127.0.0.1:8080")
}
