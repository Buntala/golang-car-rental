package request

import (
	"car-rental/entity"
	"errors"
)

type CustomerRequest struct {
	CustomerID     int    `json:"customer_id"`
	Name           string `json:"name"`
	Nik            string `json:"nik" binding:"omitempty,numeric"`
	PhoneNumber    string `json:"phone_number" binding:"omitempty,numeric"`
	MembershipName string `json:"membership_name" gorm:"-"`
}

func (ct *CustomerRequest) ToDB () entity.Customer{
	var ct_entity entity.Customer
	ct_entity.CustomerID = ct.CustomerID
	ct_entity.Name = ct.Name
	ct_entity.Nik = ct.PhoneNumber
	ct_entity.PhoneNumber = ct.PhoneNumber
	return ct_entity
}
func DBtoReqCust(ct_entity entity.Customer) CustomerRequest{
	var ct CustomerRequest
	ct.CustomerID = ct_entity.CustomerID
	ct.Name = ct_entity.Name
	ct.Nik = ct_entity.Nik
	ct.PhoneNumber = ct_entity.PhoneNumber
	return ct
}

func (ct *CustomerRequest) Validate(method string) error{
	//case method
	var err error
	switch method{
		case "get":
			err = ct.customerIDRequired()
			if err != nil {
				return err
			}
		case "post":
			err = ct.nameRequired()
			if err != nil {
				return err
			}
			err = ct.nikRequired()
			if err != nil {
				return err
			}
			err = ct.phoneNumberRequired()
			if err != nil {
				return err
			}
			err = ct.membershipRequired()
			if err != nil {
				return err
			}
		case "update":
			err = ct.customerIDRequired()
			if err != nil {
				return err
			}
			err = ct.nameRequired()
			if err != nil {
				return err
			}
			err = ct.nikRequired()
			if err != nil {
				return err
			}
			err = ct.phoneNumberRequired()
			if err != nil {
				return err
			}
			err = ct.membershipRequired()
			if err != nil {
				return err
			}
		case "delete":
			err = ct.customerIDRequired()
			if err != nil {
				return err
			}
		case "membership":
			err = ct.customerIDRequired()
			if err != nil {
				return err
			}
		default:
			return nil
	}
	return nil
}

func (ct CustomerRequest) customerIDRequired() error {
	if ct.CustomerID == 0 {
		return errors.New("customer ID is required")
	}
	return nil
}
func (ct CustomerRequest) nameRequired() error {
	if ct.Name == "" {
		return errors.New("name is required")
	}
	return nil
}
func (ct CustomerRequest) nikRequired() error {
	if ct.Nik == "" {
		return errors.New("nik is required")
	}
	return nil
}
func (ct CustomerRequest) phoneNumberRequired() error {
	if ct.PhoneNumber == "" {
		return errors.New("phone number is required")
	}
	return nil
}
func (ct CustomerRequest) membershipRequired() error {
	if ct.MembershipName == "" {
		return errors.New("membership name is required")
	}
	if !nameValidate(ct.MembershipName){
		return errors.New("membership name is invalid (Gold,Silver,Bronze only)")
	}
	return nil
}

func nameValidate(data string) bool{
	var valueObj = [3]string{"Gold","Silver","Bronze"}
	var status bool = false
	for _ , val := range valueObj{
		if val == data{
			status = true
			break;
		}
	}
	return status
}
