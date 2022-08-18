package customer

import (
	"car-rental/routes/membership"
	"errors"
	"fmt"
	"reflect"
)

type CustomerDB struct {
	CustomerID 		int    	`json:"customer_id" gorm:"primary_key"`
	Name         	string 	`json:"name"`
	Nik         	string	`json:"nik" binding:"omitempty,numeric"`
	PhoneNumber 	string	`json:"phone_number" binding:"omitempty,numeric,max=12,min=5"`
	MembershipID 	int 	`json:"membership_id,omitempty"`
	Membership 		membership.MembershipVal `gorm:"ForeignKey:MembershipId"`
	MembershipName string  `json:"membership_name" gorm:"-"`
}
func (CustomerDB) TableName() string {
	return "customer_gorm"
}
func (m *CustomerDB) Validate(method string) error{
	//case method
	var err error
	switch method{
		case "get":
			err = m.required("CustomerID")
			if err != nil {
				return err
			}
		case "post":
			err = m.required("Name")
			if err != nil {
				return err
			}
			err = m.required("Nik")
			if err != nil {
				return err
			}
			err = m.required("PhoneNumber")
			if err != nil {
				return err
			}
			err = m.required("MembershipName")
			if err != nil {
				return err
			}
		case "update":
			err = m.required("CustomerID")
			if err != nil {
				return err
			}
			err = m.required("Name")
			if err != nil {
				return err
			}
			err = m.required("Nik")
			if err != nil {
				return err
			}
			err = m.required("PhoneNumber")
			if err != nil {
				return err
			}
			err = m.required("MembershipName")
			if err != nil {
				return err
			}
		case "delete":
			err = m.required("CustomerID")
			if err != nil {
				return err
			}
		default:
			return nil
	}
	return nil
}
func (m *CustomerDB) required(column string) error {
	r := reflect.ValueOf(m)
    f := reflect.Indirect(r).FieldByName(column)
	err_str := fmt.Sprintf("%s is required", column)
	if f.String() == "" {
		return errors.New(err_str)
	}
	return nil
}

/*
type Get_Rules struct {
	MembershipId string `json:"membership_id" validate:"required,numeric"`
}

type Post_Rules struct {
	Name     string `json:"name"`
	Discount string `json:"discount" validate:"required,numeric"`
}

type Patch_Rules struct {
	MembershipId string `json:"membership_id" validate:"required,numeric"`
	Name         string `json:"name"`
	Discount     string `json:"discount" validate:"required,numeric"`
}

type Delete_Rules struct {
	MembershipId string `json:"membership_id" validate:"required,numeric"`
}*/