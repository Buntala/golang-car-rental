package main

import (
	"car-rental/router"

	"github.com/gin-gonic/gin"
)

func Router(server *gin.Engine) {
	router.Booking(server)
	router.BookingType(server)
	router.Car(server)
	router.Customer(server)
	router.Driver(server)
	router.Membership(server)
}