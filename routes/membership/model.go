package membership

import (
	"errors"
	"fmt"
	"reflect"
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
			err = m.required("MembershipId")
			if err != nil {
				return err
			}
		case "post":
			err = m.required("Name")
			if err != nil {
				return err
			}
			err = m.required("Discount")
			if err != nil {
				return err
			}
		case "update":
			err = m.required("MembershipId")
			if err != nil {
				return err
			}
			err = m.required("Name")
			if err != nil {
				return err
			}
			err = m.required("Discount")
			if err != nil {
				return err
			}
		case "delete":
			err = m.required("Membership_id")
			if err != nil {
				return err
			}
		default:
			return nil
	}
	return nil
}
func (m *MembershipVal) required(column string) error {
	r := reflect.ValueOf(m)
    f := reflect.Indirect(r).FieldByName(column)
	err_str := fmt.Sprintf("%s is required", column)
	if f.String() == "" {
		return errors.New(err_str)
	}
	return nil
}

func (m *MembershipVal) GetID() int{
	var result MembershipVal
	fmt.Printf("FROM FUNC= name:%v",m.Name)
	valid := ObjValidation(m.Name)
	if !valid{
		panic("Membership name not valid!")
	}
	conn.Where("name = ?", m.Name).First(&result)
	return result.MembershipId
}

func ObjValidation(data string) bool{
	var valueObj = [3]string{"Gold","Silver","Bronze"}
	for _ , val := range valueObj{
		if val == data{
			return false
		}
	}
	return true
}