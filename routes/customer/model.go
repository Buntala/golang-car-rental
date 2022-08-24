package customer

import (
	"car-rental/routes/membership"
	"errors"
)

type CustomerDB struct {
	CustomerID 		int    	`json:"customer_id" gorm:"primary_key"`
	Name         	string 	`json:"name"`
	Nik         	string	`json:"nik" binding:"omitempty,numeric"`
	PhoneNumber 	string	`json:"phone_number" binding:"omitempty,numeric"`
	MembershipID 	int 	`json:"membership_id,omitempty"`
	Membership 		membership.MembershipVal `gorm:"ForeignKey:MembershipID"`
	MembershipName string  `json:"membership_name" gorm:"-"`
}
func (CustomerDB) TableName() string {
	return "customer_gorm"
}
func (ct *CustomerDB) Validate(method string) error{
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
		default:
			return nil
	}
	return nil
}

func (ct CustomerDB) customerIDRequired() error {
	if ct.CustomerID == 0 {
		return errors.New("customer ID is required")
	}
	return nil
}
func (ct CustomerDB) nameRequired() error {
	if ct.Name == "" {
		return errors.New("name is required")
	}
	return nil
}
func (ct CustomerDB) nikRequired() error {
	if ct.Nik == "" {
		return errors.New("nik is required")
	}
	return nil
}
func (ct CustomerDB) phoneNumberRequired() error {
	if ct.PhoneNumber == "" {
		return errors.New("phone number is required")
	}
	return nil
}
func (ct CustomerDB) membershipRequired() error {
	if ct.MembershipName == "" {
		return errors.New("membership name is required")
	}
	return nil
}

