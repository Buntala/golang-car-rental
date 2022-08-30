package car

import (
	"errors"
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
			err = c.CarsIDRequired()
			if err != nil {
				return err
			}
		case "post":
			err = c.nameRequired()
			if err != nil {
				return err
			}
			err = c.priceRequired()
			if err != nil {
				return err
			}
			err = c.stockRequired()
			if err != nil {
				return err
			}
		case "update":
			err = c.CarsIDRequired()
			if err != nil {
				return err
			}
			err = c.nameRequired()
			if err != nil {
				return err
			}
			err = c.priceRequired()
			if err != nil {
				return err
			}
			err = c.stockRequired()
			if err != nil {
				return err
			}
		case "delete":
			err = c.CarsIDRequired()
			if err != nil {
				return err
			}
			return err
		default:
			return errors.New("validation error: CHOOSE ONE OF THESE METHODS(GET,POST,UPDATE,DELETE)")
	}
	return nil
}


func (c CarDB) CarsIDRequired() error {
	if c.CarsID == 0 {
		return errors.New("car ID is required")
	}
	return nil
}
func (c CarDB) nameRequired() error {
	if c.Name == "" {
		return errors.New("name is required")
	}
	return nil
}
func (c CarDB) priceRequired() error {
	if c.RentPriceDaily == 0 {
		return errors.New("price is required")
	}
	return nil
}
func (c CarDB) stockRequired() error {
	if c.Stock == 0 {
		return errors.New("stock is required")
	}
	return nil
}

func (m *CarDB) GetPrice() int{
	var result CarDB
	conn.First(&result,m.CarsID)
	return result.RentPriceDaily
}