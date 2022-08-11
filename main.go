package main

import (
	"car-rental/routes/driver"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	r := gin.Default()
	//r.GET("/students/get",getStudents)
	driver.Driver(r)
	r.Run(":8080")
}
