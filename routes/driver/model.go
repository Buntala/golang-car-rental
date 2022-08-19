package driver

import (
	"errors"
	"fmt"
	"reflect"
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
			err = d.required("DriverId")
			if err != nil {
				return err
			}
		case "post":
			err = d.required("Name")
			if err != nil {
				return err
			}
			err = d.required("Nik")
			if err != nil {
				return err
			}
			err = d.required("PhoneNumber")
			if err != nil {
				return err
			}
			err = d.required("DailyCost")
			if err != nil {
				return err
			}
		case "update":
			err = d.required("DriverId")
			if err != nil {
				return err
			}
			err = d.required("Name")
			if err != nil {
				return err
			}
			err = d.required("Nik")
			if err != nil {
				return err
			}
			err = d.required("PhoneNumber")
			if err != nil {
				return err
			}
			err = d.required("DailyCost")
			if err != nil {
				return err
			}
		case "delete":
			err = d.required("DriverId")
			if err != nil {
				return err
			}
		default:
			return nil
	}
	return nil
}
func (d *DriverVal) required(column string) error {
	r := reflect.ValueOf(d)
    f := reflect.Indirect(r).FieldByName(column)
	err_str := fmt.Sprintf("%s is required", column)
	if f.String() == "" {
		return errors.New(err_str)
	}
	return nil
}
/*
func (d *DriverVal) Get(id int) []Student{
	db := dbConnect()
	var result []Student
	db.Order("student_id").Find(&result)
	return result 
}*/
/*
type Get_Rules struct {
	DriverId string `json:"driver_id" db:"driver_id" validate:"required,numeric"`
}

type Post_Rules struct {
	Name        string `json:"name" db:"name" validate:"required" `
	Nik         string `json:"nik" db:"nik" validate:"required,numeric"`
	PhoneNumber string `json:"phone_number" db:"phone_number" validate:"required,numeric"`
	DailyCost   string `json:"daily_cost" db:"daily_cost" validate:"required,numeric" `
}

type Patch_Rules struct {
	DriverId    string `json:"driver_id" db:"driver_id" validate:"required,numeric"`
	Name        string `json:"name" db:"name" validate:"required"`
	Nik         string `json:"nik" db:"nik" validate:"required,numeric"`
	PhoneNumber string `json:"phone_number" db:"phone_number" validate:"required,numeric"`
	DailyCost   string `json:"daily_cost" db:"daily_cost" validate:"required,numeric"`
}

type Delete_Rules struct {
	DriverId string `json:"driver_id" db:"driver_id" validate:"required,numeric"`
}*/