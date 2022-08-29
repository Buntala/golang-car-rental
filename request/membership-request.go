package request

import (
	"car-rental/entity"
	"errors"
)

type Membership struct {
	MembershipID int    `json:"membership_id" gorm:"primary_key"`
	Name         string `json:"name"`
	Discount     int    `json:"discount" `
}

func (m *Membership) ToDB () entity.Membership{
	var m_entity entity.Membership
	m_entity.MembershipID = m.MembershipID
	m_entity.Name = m.Name
	m_entity.Discount = m.Discount
	return m_entity
}
func DBtoReqMember(m_entity entity.Membership) Membership{
	var m Membership
	m.MembershipID = m_entity.MembershipID 
	m.Name = m_entity.Name 
	m.Discount = m_entity.Discount
	return m
}
func (m *Membership) Validate(method string) error{
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

func (m Membership) MembershipIDRequired() error {
	if m.MembershipID == 0 {
		return errors.New("membership ID is required")
	}
	return nil
}
func (m Membership) nameRequired() error {
	if m.Name == "" {
		return errors.New("name is required")
	}
	return nil
}
func (m Membership) discountRequired() error {
	if m.Discount == 0 {
		return errors.New("discount is required")
	}
	return nil
}