package main

import (
	"car-rental/controller"
	"car-rental/middleware"
	"car-rental/repository"
	"car-rental/service"

	"github.com/gin-gonic/gin"
)

var (
	customerRepository repository.CustomerRepository = repository.NewCustomerRepository()
	customerService    service.CustomerService = service.NewCustomerService(customerRepository)
	customerController controller.CustomerController = controller.NewCustomer(customerService)

	membershipRepository repository.MembershipRepository = repository.NewMembershipRepository()
	membershipService    service.MembershipService = service.NewMembershipService(membershipRepository)
	membershipController controller.MembershipController = controller.NewMembership(membershipService)
	
	carRepository repository.CarRepository = repository.NewCarRepository()
	carService    service.CarService = service.NewCarService(carRepository)
	carController controller.CarController = controller.NewCar(carService)

	driverRepository repository.DriverRepository = repository.NewDriverRepository()
	driverService    service.DriverService = service.NewDriverService(driverRepository)
	driverController controller.DriverController = controller.NewDriver(driverService)

	bookingTypeRepository repository.BookingTypeRepository = repository.NewBookingTypeRepository()
	bookingTypeService    service.BookingTypeService = service.NewBookingTypeService(bookingTypeRepository)
	bookingTypeController controller.BookingTypeController = controller.NewBookingType(bookingTypeService)

	bookingRepository repository.BookingRepository = repository.NewBookingRepository()
	bookingService    service.BookingService = service.NewBookingService(bookingRepository)
	bookingController controller.BookingController = controller.NewBooking(bookingService)

)

func main() {
	server := gin.New()
	server.Use(gin.Logger())
	/*server.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))*/
	//server.Use(gin.Recovery())
	server.Use(gin.CustomRecovery(middleware.ErrorResponse))
	cust_r := server.Group("/customers")
	cust_r.GET("/", customerController.FindAll)
	cust_r.GET("/:id", customerController.FindOne)
	cust_r.POST("/", customerController.Save)
	cust_r.PATCH("/:id",  customerController.Update)
	cust_r.DELETE("/:id", customerController.Delete)
	cust_r.POST("/:id/membership", customerController.SaveMembership)

	
	member_r := server.Group("/memberships")

	member_r.GET("/", membershipController.FindAll)
	member_r.GET("/:id", membershipController.FindOne)
	member_r.POST("/",  membershipController.Save)
	member_r.PATCH("/:id",  membershipController.Update)
	member_r.DELETE("/:id",  membershipController.Delete)

	car_r := server.Group("/cars")

	car_r.GET("/", carController.FindAll)
	car_r.GET("/:id", carController.FindOne)
	car_r.POST("/",  carController.Save)
	car_r.PATCH("/:id",  carController.Update)
	car_r.DELETE("/:id",  carController.Delete)

	driver_r := server.Group("/driver")

	driver_r.GET("/", driverController.FindAll)
	driver_r.GET("/:id", driverController.FindOne)
	driver_r.POST("/",  driverController.Save)
	driver_r.PATCH("/:id",  driverController.Update)
	driver_r.DELETE("/:id",  driverController.Delete)
	
	bookingtype_r := server.Group("/booking-type")

	bookingtype_r.GET("/", bookingTypeController.FindAll)
	bookingtype_r.GET("/:id", bookingTypeController.FindOne)
	bookingtype_r.POST("/",  bookingTypeController.Save)
	bookingtype_r.PATCH("/:id",  bookingTypeController.Update)
	bookingtype_r.DELETE("/:id",  bookingTypeController.Delete)

	booking_r := server.Group("/bookings")

	booking_r.GET("/", bookingController.FindAll)
	booking_r.GET("/:id", bookingController.FindOne)
	booking_r.POST("/",  bookingController.Save)
	booking_r.PATCH("/:id",  bookingController.Update)
	booking_r.DELETE("/:id",  bookingController.Delete)
	booking_r.POST("/:id/finish",  bookingController.SaveExtend)
	booking_r.POST("/:id/extend",  bookingController.SaveFinished)

	server.Run("127.0.0.1:8080")
}

