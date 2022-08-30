package membership

import (
	"errors"
)

type MembershipVal struct {
	MembershipID int    `json:"membership_id" gorm:"primary_key"`
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
			err = m.MembershipIDRequired()
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
			err = m.MembershipIDRequired()
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
			err = m.MembershipIDRequired()
			if err != nil {
				return err
			}
		default:
			return nil
	}
	return nil
}

func (m MembershipVal) MembershipIDRequired() error {
	if m.MembershipID == 0 {
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
func (m *MembershipVal) GetID() (int,error){
	var result MembershipVal
	valid := objValidation(m.Name)
	if !valid{
		return 0,errors.New("membership name not valid! (Gold,Silver,Bronze only)")
	}
	conn.Where("name = ?", m.Name).First(&result)
	return result.MembershipID,nil
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
func (m *MembershipVal) GetDiscount() int{
	var result MembershipVal
	conn.First(&result,m.MembershipID)
	return result.Discount
}
func (m *MembershipVal) GetName() string{	
	var result MembershipVal
	conn.First(&result,m.MembershipID)
	return result.Name
}
