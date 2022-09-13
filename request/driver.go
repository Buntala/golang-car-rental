package request

import (
	"car-rental/entity"
	"errors"
)

type Driver struct {
	DriverID    int    `json:"driver_id"  binding:"numeric"`
	Name        string `json:"name" `
	Nik         string `json:"nik" binding:"omitempty,numeric"`
	PhoneNumber string `json:"phone_number" binding:"omitempty,numeric,max=12,min=10"`
	DailyCost   int    `json:"daily_cost"`
}

func (d *Driver) ToDB () entity.Driver{
	var d_entity entity.Driver
	d_entity.DriverID = d.DriverID
	d_entity.Name = d.Name
	d_entity.Nik = d.Nik
	d_entity.PhoneNumber = d.PhoneNumber
	d_entity.DailyCost = d.DailyCost

	return d_entity
}
func DBtoReqDriver(d_entity entity.Driver) Driver{
	var d Driver
	d.DriverID = d_entity.DriverID 
	d.Name = d_entity.Name 
	d.Nik =	d_entity.Nik
	d.PhoneNumber = d_entity.PhoneNumber
	d.DailyCost = d_entity.DailyCost
	return d
}

func (d *Driver) Validate(method string) error {
	//case method
	var err error
	switch method {
	case "get":
		err = d.driverIDRequired()
		if err != nil {
			return err
		}
	case "post":
		err = d.nameRequired()
		if err != nil {
			return err
		}
		err = d.nikRequired()
		if err != nil {
			return err
		}
		err = d.phoneNumberRequired()
		if err != nil {
			return err
		}
		err = d.dailyCostRequired()
		if err != nil {
			return err
		}
	case "update":
		err = d.driverIDRequired()
		if err != nil {
			return err
		}
		err = d.nameRequired()
		if err != nil {
			return err
		}
		err = d.nikRequired()
		if err != nil {
			return err
		}
		err = d.phoneNumberRequired()
		if err != nil {
			return err
		}
		err = d.dailyCostRequired()
		if err != nil {
			return err
		}
	case "delete":
		err = d.driverIDRequired()
		if err != nil {
			return err
		}
	default:
		return nil
	}
	return nil
}

func (d Driver) driverIDRequired() error {
	if d.DriverID == 0 {
		return errors.New("driver ID is required")
	}
	return nil
}
func (d Driver) nameRequired() error {
	if d.Name == "" {
		return errors.New("name is required")
	}
	return nil
}
func (d Driver) nikRequired() error {
	if d.Nik == "" {
		return errors.New("nik is required")
	}
	return nil
}
func (d Driver) phoneNumberRequired() error {
	if d.PhoneNumber == "" {
		return errors.New("phoneNumber is required")
	}
	return nil
}
func (d Driver) dailyCostRequired() error {
	if d.DailyCost == 0 {
		return errors.New("daily cost is required")
	}
	return nil
}
