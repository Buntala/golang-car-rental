package request

import (
	"car-rental/entity"
	"errors"
	"time"
)

type Booking struct {
	BookingID       int        `json:"booking_id"`
	CustomerID      int        `json:"customer_id"`
	CarsID          int        `json:"car_id"`
	StartTime       time.Time  `json:"start_time" binding:"omitempty"`
	EndTime         time.Time  `json:"end_time" binding:"omitempty"`
	Finished        bool       `json:"finished"`
	BookingTypeName string     `json:"booking_type"`
	BookingTypeID   int        `json:"booking_type_id"`
	DriverID        int        `json:"driver_id,omitempty"`
}

func (Booking) TableName() string {
	return "booking_table"
}

func (b *Booking) ToDB () entity.Booking{
	var b_entity entity.Booking
	b_entity.BookingID = b.BookingID
	b_entity.CustomerID = b.CustomerID
	b_entity.CarsID = b.CarsID
	b_entity.StartTime = b.StartTime
	b_entity.EndTime = b.EndTime
	b_entity.TotalCost = 0 //b.TotalCost
	b_entity.Finished = b.Finished
	b_entity.Discount = 0 //b.Discount
	b_entity.BookingTypeID = 0 // b.BookingTypeName
	b_entity.DriverID = b.DriverID
	b_entity.TotalDriverCost = 0 //
	b_entity.DriverIncentive = 0 //
	return b_entity
}
func DBtoReqBooking(b_entity entity.Booking) Booking{
	var b Booking
	b.BookingID = b_entity.BookingID

	b.CustomerID = b_entity.CustomerID

	b.CarsID = b_entity.CarsID
	
	b.StartTime = b_entity.StartTime

	b.EndTime = b_entity.EndTime 

	b.Finished = b_entity.Finished 


	b.DriverID = b_entity.DriverID 

	return b
}


func (b *Booking) Validate(method string) error{
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
		default:
			return errors.New("validation error: CHOOSE ONE OF THESE METHODS(GET,POST,UPDATE,DELETE,FINISHED,CANCEL,EXTEND)")
	}
	return nil
}

func (b Booking) bookingIDRequired() error {
	if b.BookingID == 0 {
		return errors.New("booking ID is required")
	}
	return nil
}
func (b Booking) carsIDRequired() error {
	if b.CarsID == 0 {
		return errors.New("car ID is required")
	}
	return nil
}
func (b Booking) customerIDRequired() error {
	if b.CustomerID == 0 {
		return errors.New("customer ID is required")
	}
	return nil
}
func (b Booking) startTimeRequired() error {
	if b.StartTime.IsZero() {
		return errors.New("start time is required")
	}
	return nil
}
func (b Booking) EndTimeRequired() error {
	if b.EndTime.IsZero() {
		return errors.New("start time is required")
	}
	return nil
}
func (b Booking) EndTimeLater() error {
	if b.StartTime.After(b.EndTime) {
		return errors.New("end time has to be later than start time")
	}
	return nil
}
func (b Booking) bookingTypeRequired() error {
	if b.BookingTypeName == "" {
		return errors.New("booking type is required")
	}
	if !bookingTypeValidate(b.BookingTypeName){
		return errors.New("membership name is invalid (Gold,Silver,Bronze only)")
	}
	if b.BookingTypeName == "Car & Driver"{
		err := b.driverIDRequired()
		return err
	}
	return nil
}
func (b Booking) driverIDRequired() error {
	if b.DriverID == 0 {
		return errors.New("driver ID is required")
	}
	return nil
}


func bookingTypeValidate(data string) bool{
	var valueObj = [3]string{"Gold","Silver","Bronze"}
	var status bool = false
	for _ , val := range valueObj{
		if val == data{
			status = true
			break;
		}
	}
	return status
}