package customer

import (
	"car-rental/routes/membership"
	"car-rental/utilities/db"
	"car-rental/utilities/responseHandler"

	//"fmt"
	//"github.com/jmoiron/sqlx"

	"gorm.io/gorm"
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

func DBGetCustomerOne(params CustomerDB) CustomerDB{
	var result CustomerDB
	err := conn.First(&result,params.CustomerID).Error
	responseHandler.ErrorHandler(err)
	return result
}

func DBInsertCustomer(params *CustomerDB) {
	var member_strt membership.MembershipVal 
	member_strt.Name= params.MembershipName
	params.MembershipID = member_strt.GetID()
	err := conn.Create(&params).Error
	responseHandler.ErrorHandler(err)
}
func DBUpdateCustomer(params *CustomerDB) {
	var member_strt membership.MembershipVal 
	member_strt.Name= params.MembershipName
	params.MembershipID = member_strt.GetID()
	err := conn.Updates(&params).Error
	responseHandler.ErrorHandler(err)
}

func DBDeleteCustomer(params *CustomerDB) {
	err := conn.Delete(&params).Error
	responseHandler.ErrorHandler(err)
}