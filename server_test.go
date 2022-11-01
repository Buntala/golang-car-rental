package main

import (
	"car-rental/middleware"
	"car-rental/router"
	"fmt"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func routing(server *gin.Engine) {
	router.Booking(server)
	router.BookingType(server)
	router.Car(server)
	router.Customer(server)
	router.Driver(server)
	router.Membership(server)
}

func TestHttpRequest(t *testing.T) {
	server := gin.New()
	server.Use(gin.Logger())
	server.Use(gin.CustomRecovery(middleware.ErrorResponse))

	routing(server)
	req := httptest.NewRequest("POST", "/cars/", strings.NewReader(`{"name":"Honda H-RV","stock": 1,"rent_price_daily":20000}")`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
	if resp.StatusCode != 200 {
		t.Fatal("Status Code Not OK")
	}
}
