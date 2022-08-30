package main

import (
	"car-rental/routes/booking"
	bookingType "car-rental/routes/booking-type"
	"car-rental/routes/car"
	"car-rental/routes/customer"
	"car-rental/routes/driver"
	"car-rental/routes/membership"
	"car-rental/utilities/db"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	//godotenv.Load()
	fmt.Println(os.Getenv("PGHOST"))
	r := gin.New()
	r.Use(gin.Logger(),gin.Recovery())
	var conn *gorm.DB = db.DbConnectGorm()
	conn.AutoMigrate(&booking.BookingDB{})
	car.Car(r)
	driver.Driver(r)
	membership.Membership(r)
	customer.Customer(r)
	booking.Booking(r)
	bookingType.BookingType(r)
	r.Run(":8080")
}
