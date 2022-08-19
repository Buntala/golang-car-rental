package car

import (
	"errors"
	"fmt"
	"reflect"
)
type CarDB struct {
	CarsID    			int    	`json:"car_id"  binding:"numeric" gorm:"primary_key"`
	Name        		string 	`json:"name" `
	RentPriceDaily   	int    	`json:"rent_price_daily" `
	Stock 				int 	`json:"stock" `
}
func (CarDB) TableName() string {
	return "cars"
}

func (c *CarDB) Validate(method string) error{
	//case method
	var err error
	switch method{
		case "get":	
			return c.CarsIDReqired()
		case "post":
			err = c.nameReqired()
			if err != nil {
				return err
			}
			err = c.priceReqired()
			if err != nil {
				return err
			}
			err = c.stockRequired()
			if err != nil {
				return err
			}
		case "update":
			err = c.CarsIDReqired()
			if err != nil {
				return err
			}
			err = c.nameReqired()
			if err != nil {
				return err
			}
			err = c.priceReqired()
			if err != nil {
				return err
			}
			err = c.stockRequired()
			if err != nil {
				return err
			}
		case "delete":
			return c.CarsIDReqired()
		default:
			return errors.New("CHOOSE ONE OF THESE METHODS(GET,POST,UPDATE,DELETE)")
	}
	return nil
}


func (c CarDB) CarsIDReqired() error {
	err := c.required("CarsID")
	return err
}
func (c CarDB) nameReqired() error {
	err := c.required("Name")
	return err
}
func (c CarDB) priceReqired() error {
	err := c.required("RentPriceDaily")
	return err
}
func (c CarDB) stockRequired() error {
	err := c.required("Stock")
	return err
}

func (c *CarDB) required(column string) error {
	r := reflect.ValueOf(c)
    f := reflect.Indirect(r).FieldByName(column)
	err_str := fmt.Sprintf("%s is required", column)
	if f.String() == "" {
		return errors.New(err_str)
	}
	return nil
}