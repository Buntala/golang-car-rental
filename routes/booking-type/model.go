package bookingType

import (
	"errors"
)
type BookingTypeDB struct {
	BookingTypeID    int    	`json:"booking_type_id"  binding:"numeric" gorm:"primary_key"`
	BookingType      string 	`json:"booking_type" `
	Description	   	 string    	`json:"description" `
}
func (BookingTypeDB) TableName() string {
	return "booking_type"
}

func (bt *BookingTypeDB) Validate(method string) error{
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


func (bt BookingTypeDB) bookingTypeIDRequired() error {
	if bt.BookingTypeID == 0 {
		return errors.New("booking type ID is required")
	}
	return nil
}
func (bt BookingTypeDB) bookingTypeRequired() error {
	if bt.BookingType == "" {
		return errors.New("booking type is required")
	}
	return nil
}
func (bt BookingTypeDB) desciptionRequired() error {
	if bt.Description == "" {
		return errors.New("description is required")
	}
	return nil
}

//add error return
func (bt BookingTypeDB) GetID() (BookingTypeDB,error) {
	var result BookingTypeDB
	if err:= objValidation(bt.BookingType);err!=nil{
		return result, err
	}
	conn.Where("booking_type = ?", bt.BookingType).First(&result)
	return result,nil
}


func objValidation(data string) error{
	var valueObj = [2]string{"Car Only","Car & Driver"}
	for _ , val := range valueObj{
		if val == data{
			return nil
		}
	}
	return errors.New("invalid booking type")
}