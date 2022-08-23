package membership

import (
	"errors"
)

type MembershipVal struct {
	MembershipId int    `json:"membership_id" gorm:"primary_key"`
	Name         string `json:"name"`
	Discount     int    `json:"discount" `
}
func (MembershipVal) TableName() string {
	return "membership"
}
func (m *MembershipVal) Validate(method string) error{
	//case method
	var err error
	switch method{
		case "get":
			err = m.membershipIDRequired()
			if err != nil {
				return err
			}
		case "post":
			err = m.nameRequired()
			if err != nil {
				return err
			}
			err = m.discountRequired()
			if err != nil {
				return err
			}
		case "update":
			err = m.membershipIDRequired()
			if err != nil {
				return err
			}
			err = m.nameRequired()
			if err != nil {
				return err
			}
			err = m.discountRequired()
			if err != nil {
				return err
			}
		case "delete":
			err = m.membershipIDRequired()
			if err != nil {
				return err
			}
		default:
			return nil
	}
	return nil
}

func (m MembershipVal) membershipIDRequired() error {
	if m.MembershipId == 0 {
		return errors.New("membership ID is required")
	}
	return nil
}
func (m MembershipVal) nameRequired() error {
	if m.Name == "" {
		return errors.New("name is required")
	}
	return nil
}
func (m MembershipVal) discountRequired() error {
	if m.Discount == 0 {
		return errors.New("discount is required")
	}
	return nil
}
func (m *MembershipVal) GetID() int{
	var result MembershipVal
	valid := objValidation(m.Name)
	if !valid{
		panic("Membership name not valid!")
	}
	conn.Where("name = ?", m.Name).First(&result)
	return result.MembershipId
}

func objValidation(data string) bool{
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