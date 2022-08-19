package main

import (
	"car-rental/routes/customer"
	"car-rental/routes/driver"
	"car-rental/routes/membership"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	//godotenv.Load()
	fmt.Println(os.Getenv("PGHOST"))
	r := gin.Default()
	//r.GET("/students/get",getStudents)
	driver.Driver(r)
	membership.Membership(r)
	customer.Customer(r)
	r.Run(":8080")
}
