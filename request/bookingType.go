package request

import (
	"car-rental/entity"
	"errors"
)

type BookingType struct {
	BookingTypeID    int    	`json:"booking_type_id"  binding:"numeric"`
	BookingType      string 	`json:"booking_type" `
	Description	   	 string    	`json:"description" `
}

func (bt *BookingType) ToDB () entity.BookingType{
	var bt_entity entity.BookingType
	bt_entity.BookingTypeID = bt.BookingTypeID
	bt_entity.BookingType = bt.BookingType
	bt_entity.Description = bt.Description
	return bt_entity
}
func DBtoReqBookingType(bt_entity entity.BookingType) BookingType{
	var bt BookingType
	bt.BookingTypeID = bt_entity.BookingTypeID 
	bt.BookingType = bt_entity.BookingType 
	bt.Description = bt_entity.Description
	return bt
}
func (bt *BookingType) Validate(method string) error{
	//case method
	var err error
	switch method{
		case "get":	
			return bt.bookingTypeIDRequired()
		case "post":
			err = bt.bookingTypeRequired()
			if err != nil {
				return err
			}
			err = bt.desciptionRequired()
			if err != nil {
				return err
			}
		case "update":
			err = bt.bookingTypeIDRequired()
			if err != nil {
				return err
			}
			err = bt.bookingTypeRequired()
			if err != nil {
				return err
			}
			err = bt.desciptionRequired()
			if err != nil {
				return err
			}
		case "delete":
			return bt.bookingTypeIDRequired()
		default:
			return errors.New("CHOOSE ONE OF THESE METHODS(GET,POST,UPDATE,DELETE)")
	}
	return nil
}

func (bt BookingType) bookingTypeIDRequired() error {
	if bt.BookingTypeID == 0 {
		return errors.New("bookingType ID is required")
	}
	return nil
}
func (bt BookingType) bookingTypeRequired() error {
	if bt.BookingType == "" {
		return errors.New("booking type is required")
	}
	return nil
}
func (bt BookingType) desciptionRequired() error {
	if bt.Description == "" {
		return errors.New("description is required")
	}
	return nil
}
