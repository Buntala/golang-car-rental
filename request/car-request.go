package request

import (
	"car-rental/entity"
	"errors"

	"github.com/go-playground/validator/v10"
)

type Car struct {
	CarID    			int    	`json:"car_id"  binding:"numeric"`
	Name        		string 	`json:"name" `
	RentPriceDaily   	int    	`json:"rent_price_daily" `
	Stock 				int 	`json:"stock" `
}

func (c *Car) ToDB () entity.Car{
	var c_entity entity.Car
	c_entity.CarsID = c.CarID
	c_entity.Name = c.Name
	c_entity.RentPriceDaily = c.RentPriceDaily
	c_entity.Stock = c.Stock
	return c_entity
}
func DBtoReqCar(c_entity entity.Car) Car{
	var c Car
	c.CarID = c_entity.CarsID 
	c.Name = c_entity.Name 
	c.RentPriceDaily = c_entity.RentPriceDaily
	c.Stock = c_entity.Stock
	return c
}
func (c *Car) Validate(method string) error{
	//case method
	var err error
	switch method{
		case "get":
			err = c.CarIDRequired()
			if err != nil {
				return err
			}
		case "post":
			validator := validator.New()
			validator.Struct(c)
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
			validator := validator.New()
			validator.Struct(c)
			err = c.CarIDRequired()
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
			err = c.CarIDRequired()
			if err != nil {
				return err
			}
		default:
			return nil
	}
	return nil
}

func (c Car) CarIDRequired() error {
	if c.CarID == 0 {
		return errors.New("car ID is required")
	}
	return nil
}
func (c Car) nameRequired() error {
	if c.Name == "" {
		return errors.New("name is required")
	}
	return nil
}
func (c Car) priceRequired() error {
	if c.RentPriceDaily == 0 {
		return errors.New("price is required")
	}
	return nil
}
func (c Car) stockRequired() error {
	if c.Stock == 0 {
		return errors.New("stock ID is required")
	}
	return nil
}