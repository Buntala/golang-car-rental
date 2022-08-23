package driver

import (
	"errors"
)
type DriverVal struct {
	DriverId    int    `json:"driver_id"  binding:"numeric" gorm:"primary_key"`
	Name        string `json:"name" `
	Nik         string `json:"nik" binding:"omitempty,numeric"`
	PhoneNumber string `json:"phone_number" binding:"omitempty,numeric,max=12,min=5"`
	DailyCost   int    `json:"daily_cost" `
}
func (DriverVal) TableName() string {
	return "driver"
}

func (d *DriverVal) Validate(method string) error{
	//case method
	var err error
	switch method{
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

func (d DriverVal) driverIDRequired() error {
	if d.DriverId == 0 {
		return errors.New("driver ID is required")
	}
	return nil
}
func (d DriverVal) nameRequired() error {
	if d.Name == "" {
		return errors.New("name is required")
	}
	return nil
}
func (d DriverVal) nikRequired() error {
	if d.Nik == "" {
		return errors.New("nik is required")
	}
	return nil
}
func (d DriverVal) phoneNumberRequired() error {
	if d.PhoneNumber == "" {
		return errors.New("phoneNumber is required")
	}
	return nil
}
func (d DriverVal) dailyCostRequired() error {
	if d.DailyCost == 0 {
		return errors.New("daily cost is required")
	}
	return nil
}

func (d *DriverVal) GetCost() int {
	var result DriverVal
	conn.First(&result,d.DriverId)
	return result.DailyCost
}
