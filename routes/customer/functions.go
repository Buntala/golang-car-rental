package customer

import (
	"car-rental/routes/membership"
	"car-rental/utilities/db"
	"errors"

	//"fmt"
	//"github.com/jmoiron/sqlx"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)
var(
	conn *gorm.DB = db.DbConnectGorm()
	//conn *sqlx.DB = db.DbConnect()
)
func DBGetCustomerAll() []CustomerDB{
	conn.AutoMigrate(&CustomerDB{})
	var result []CustomerDB
	conn.Order("customer_id desc").Find(&result)
	return result
}

func DBGetCustomerOne(params CustomerDB) (CustomerDB,error){
	var result CustomerDB
	err := conn.First(&result,params.CustomerID).Error
	return result,err
}

func DBInsertCustomer(params *CustomerDB) error{
	var member_strt membership.MembershipVal 
	member_strt.Name= params.MembershipName
	params.MembershipID = member_strt.GetID()
	err := conn.Create(&params).Error
	return err
}
func DBUpdateCustomer(params *CustomerDB) error {
	var member_strt membership.MembershipVal 
	member_strt.Name= params.MembershipName
	params.MembershipID = member_strt.GetID()
	status := conn.Updates(&params)
	if err := status.Error; err!=nil{
		return err
	}
	if status.RowsAffected == 0{
		return errors.New("no data with the input id")
	}
	//responseHandler.ErrorHandler(err)
	return nil
}

func DBDeleteCustomer(params *CustomerDB) error{
	status := conn.Clauses(clause.Returning{}).Delete(&params)
	if err := status.Error; err!=nil{
		return err
	}
	if status.RowsAffected == 0{
		return errors.New("no data with the input id")
	}
	return nil
}

func DBUpdateMembershipStatus(params *CustomerDB) error{
	var member_strt membership.MembershipVal 
	member_strt.Name= params.MembershipName
	params.MembershipID = member_strt.GetID()
	status := conn.Model(&params).Update("membership_id",params.MembershipID)
	if err:= status.Error;err!=nil{
		return err
	}	
	if status.RowsAffected == 0{
		return errors.New("no data with the input id")
	}
	return nil
}
