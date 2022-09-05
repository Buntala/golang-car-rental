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
	StartTime       string     `json:"start_time" binding:"omitempty"`
	EndTime         string  `json:"end_time" binding:"omitempty"`
	TotalCost       int        `json:"total_cost"`
	Finished        bool       `json:"finished"`
	Discount        int               `json:"discount"`
	BookingTypeName string     `json:"booking_type"`
	//BookingTypeID   int        `json:"booking_type_id"`
	DriverID        int        `json:"driver_id,omitempty"`
	TotalDriverCost int               `json:"total_driver_cost"`
	DriverIncentive int               `json:"driver_incentive"`
}

func (Booking) TableName() string {
	return "booking_table"
}

func (b *Booking) ToDB () entity.Booking{
	var b_entity entity.Booking
	b_entity.BookingID = b.BookingID
	b_entity.CustomerID = b.CustomerID 
	b_entity.CarsID = b.CarsID
	b_entity.StartTime,_ = time.Parse("2006-01-02",b.StartTime)
	b_entity.EndTime,_ = time.Parse("2006-01-02",b.EndTime)
	b_entity.TotalCost = 0 
	b_entity.Finished = b.Finished
	b_entity.Discount = 0 
	b_entity.BookingTypeID = 0 
	b_entity.DriverID = b.DriverID
	b_entity.TotalDriverCost = 0 
	b_entity.DriverIncentive = 0 
	return b_entity
}
func DBtoReqBooking(b_entity entity.Booking) Booking{
	var b Booking
	b.BookingID = b_entity.BookingID

	b.CustomerID = b_entity.CustomerID

	b.CarsID = b_entity.CarsID
	
	b.StartTime = b_entity.StartTime.Format("2006-01-02")

	b.EndTime = b_entity.EndTime.Format("2006-01-02")

	b.TotalCost = b_entity.TotalCost
	
	b.Finished = b_entity.Finished 

	b.Discount = b_entity.Discount

	b.DriverID = b_entity.DriverID 

	b.TotalDriverCost = b_entity.TotalDriverCost

	b.DriverIncentive = b_entity.DriverIncentive
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
	if b.StartTime == "" {
		return errors.New("start time is required")
	}
	_, err := time.Parse("2006-01-02",b.StartTime)
	return err
}
func (b Booking) EndTimeRequired() error {
	if b.EndTime == "" {
		return errors.New("start time is required")
	}
	_, err := time.Parse("2006-01-02",b.EndTime)
	return err
}
func (b Booking) EndTimeLater() error {
	start,_ := time.Parse("2006-01-02",b.StartTime)
	end,_ := time.Parse("2006-01-02",b.EndTime)
	if start.After(end) {
		return errors.New("end time has to be later than start time")
	}
	return nil
}
func (b Booking) bookingTypeRequired() error {
	if b.BookingTypeName == "" {
		return errors.New("booking type is required")
	}
	if !bookingTypeValidate(b.BookingTypeName){
		return errors.New("booking type is invalid (Car Only,Car and Driver)")
	}
	if b.BookingTypeName == "Car & Driver"{
		err := b.driverIDRequired()
		return err
	}
	err := b.driverIDNotAllowed()
	return err
}
func (b Booking) driverIDRequired() error {
	if b.DriverID == 0 {
		return errors.New("driver ID is required")
	}
	return nil
}
func (b Booking) driverIDNotAllowed() error {
	if b.DriverID != 0 {
		return errors.New("cant insert driver id on car only booking")
	}
	return nil
}

func bookingTypeValidate(data string) bool{
	var valueObj = [2]string{"Car Only","Car & Driver"}
	var status bool = false
	for _ , val := range valueObj{
		if val == data{
			status = true
			break;
		}
	}
	return status
}