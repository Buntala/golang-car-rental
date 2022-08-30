package booking

import (
	bookingType "car-rental/routes/booking-type"
	"car-rental/routes/car"
	"car-rental/routes/customer"
	"car-rental/routes/driver"
	"car-rental/routes/membership"
	"errors"
	"time"

	"gorm.io/gorm"
)

type BookingDB struct {
	BookingID 		int    	`json:"booking_id" gorm:"primary_key"`
	CustomerID      int    	`json:"customer_id"`
	Customer		customer.CustomerDB `json:"-" gorm:"foreignKey:CustomerID"`
	CarsID			int 	`json:"car_id"`
	Cars			car.CarDB `json:"-" gorm:"foreignKey:CarsID"`
	StartTime       time.Time	`json:"start_time" binding:"omitempty"`
	EndTime 		time.Time	`json:"end_time" binding:"omitempty"`
	TotalCost 		int 	`json:"total_cost"`
	Finished 		bool	`json:"finished"`
	Discount		int 	`json:"discount"`
	BookingTypeName string 	`json:"booking_type"`
	BookingTypeID	int 	`json:"booking_type_id"`
	BookingType		bookingType.BookingTypeDB `json:"-" gorm:"foreignKey:BookingTypeID"`
	DriverID		int		`json:"driver_id,omitempty"`
	Driver			driver.DriverVal `json:"-" gorm:"foreignKey:DriverID"` 
	TotalDriverCost int		`json:"total_driver_cost"`
	DriverIncentive int		`json:"driver_incentive"`
	Deleted 		gorm.DeletedAt 
}
func (BookingDB) TableName() string {
	return "booking_table"
}

func (b *BookingDB) Validate(method string) error{
	//case method
	var err error
	switch method{
		case "get":
			err = b.bookingIDRequired()
			if err != nil {
				return err
			}
		case "post":
			err = b.carsIDRequired()
			if err != nil {
				return err
			}
			err = b.customerIDRequired()
			if err != nil {
				return err
			}
			err = b.startTimeRequired()
			if err != nil {
				return err
			}
			err = b.EndTimeRequired()
			if err != nil {
				return err
			}
			err = b.EndTimeLater()
			if err != nil {
				return err
			}
			err = b.bookingTypeRequired()
			if err != nil {
				return err
			}
		case "update":
			err = b.bookingIDRequired()
			if err != nil {
				return err
			}
			err = b.carsIDRequired()
			if err != nil {
				return err
			}
			err = b.bookingTypeRequired()
			if err != nil {
				return err
			}
		case "delete":
			err = b.bookingIDRequired()
			if err != nil {
				return err
			}
		case "extend":
			err = b.bookingIDRequired()
			if err != nil {
				return err
			}
			err = b.EndTimeRequired()
			if err != nil {
				return err
			}
		case "finish":
			err = b.bookingIDRequired()
			if err != nil {
				return err
			}
		case "cancel":
			err = b.bookingIDRequired()
			if err != nil {
				return err
			}
		default:
			return errors.New("validation error: CHOOSE ONE OF THESE METHODS(GET,POST,UPDATE,DELETE,FINISHED,CANCEL,EXTEND)")
	}
	return nil
}

func (b BookingDB) bookingIDRequired() error {
	if b.BookingID == 0 {
		return errors.New("booking ID is required")
	}
	return nil
}
func (b BookingDB) carsIDRequired() error {
	if b.CarsID == 0 {
		return errors.New("car ID is required")
	}
	return nil
}
func (b BookingDB) customerIDRequired() error {
	if b.CustomerID == 0 {
		return errors.New("customer ID is required")
	}
	return nil
}
func (b BookingDB) startTimeRequired() error {
	if b.StartTime.IsZero() {
		return errors.New("start time is required")
	}
	return nil
}
func (b BookingDB) EndTimeRequired() error {
	if b.EndTime.IsZero() {
		return errors.New("start time is required")
	}
	return nil
}
func (b BookingDB) EndTimeLater() error {
	if b.StartTime.After(b.EndTime) {
		return errors.New("end time has to be later than start time")
	}
	return nil
}
func (b BookingDB) bookingTypeRequired() error {
	if b.BookingTypeName == "" {
		return errors.New("booking type is required")
	}
	if b.BookingTypeName == "Car & Driver"{
		err := b.driverIDRequired()
		return err
	}
	return nil
}
func (b BookingDB) driverIDRequired() error {
	if b.DriverID == 0 {
		return errors.New("driver ID is required")
	}
	return nil
}


func (b *BookingDB) calculate() error{
	var bookType bookingType.BookingTypeDB
	bookType.BookingType = b.BookingTypeName
	bookType, err := bookType.GetID()
	if err!=nil{
		return err
	}
	b.BookingTypeID = bookType.BookingTypeID
	//total cost
	duration := int(b.EndTime.Sub(b.StartTime).Hours()/24) + 1
	var car car.CarDB
	car.CarsID= b.CarsID
	b.TotalCost = duration * car.GetPrice()
	//discount
	var cust customer.CustomerDB
	cust.CustomerID = b.CustomerID
	res,err :=customer.DBGetCustomerOne(cust)
	if err!=nil{
		return errors.New("customer id is invalid")
	}
	if res.MembershipID !=0 {
		var member membership.MembershipVal
		member.MembershipID = res.MembershipID
		b.Discount = b.TotalCost * member.GetDiscount() / 100
	}
	if (b.BookingTypeID == 2){
		var driver driver.DriverVal
		driver.DriverId = b.DriverID
		driverCost,err := driver.GetCost()
		if err !=nil{
			return err
		}
		b.TotalDriverCost = duration * driverCost
		b.DriverIncentive = int(float64(b.TotalCost) * 0.05)
	}
	return nil
}
func (b *BookingDB) availabilityCheck() error{
	var car car.CarDB
	var carBooked []BookingDB
	success := conn.Find(&car,b.CarsID).RowsAffected
	if success == 0{
		return errors.New("car id is invalid")
	}
	stock := car.Stock
	booked:= conn.Where("cars_id= ?",b.CarsID).Where("booking_id != ?",b.BookingID).Where(
				conn.Where(conn.Where(
				"start_time >= ?",b.StartTime).Where(
				"start_time <= ?",b.EndTime)).Or(conn.Where(
				"end_time >= ?",b.StartTime).Where(
				"end_time <= ?",b.EndTime)).Or(conn.Where(
				"start_time <= ?" , b.StartTime).Where(
				"end_time >= ?" , b.EndTime))).Find(&carBooked).RowsAffected
	if int64(stock) <= booked{
		return errors.New("car is fully booked")
	}
	//driver availability
	if b.BookingTypeName >= "Car & Driver"{
		var driverBook BookingDB
		booked:= conn.Where("driver_id= ?",b.DriverID).Where("booking_id != ?",b.BookingID).Where(conn.Where(conn.Where(
			"start_time >= ?",b.StartTime).Where(
			"start_time <= ?",b.EndTime)).Or(conn.Where(
			"end_time >= ?",b.StartTime).Where(
			"end_time <= ?",b.EndTime)).Or(conn.Where(
			"start_time <= ?" , b.StartTime).Where(
			"end_time >= ?" , b.EndTime))).Find(&driverBook).RowsAffected
		if booked >= 1{
			return errors.New("driver is booked")
		}
	}
	return nil
}