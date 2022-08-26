package customer

import (
	"car-rental/routes/membership"
	"car-rental/utilities/db"
	"errors"

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
	for i := range result{
		if result[i].MembershipID != 0{
			result[i].fillMembershipName()
		}
	}
	return result
}

func DBGetCustomerOne(params CustomerDB) (CustomerDB,error){
	var result CustomerDB
	err := conn.First(&result,params.CustomerID).Error
	if params.MembershipID == 0{
		return result,err
	}
	result.fillMembershipName()
	return result,err
}

func DBInsertCustomer(params *CustomerDB) error{
	if params.MembershipName == ""{
		if err := conn.Omit("membership_id").Create(&params).Error; err != nil{
			return err	
		}
		return nil
	}
	var member_strt membership.MembershipVal 
	member_strt.Name= params.MembershipName
	membershipID,err := member_strt.GetID()
	if err !=nil{
		return err
	}
	params.MembershipID = membershipID
	if err := conn.Create(&params).Error; err != nil{
		return err	
	}
	params.MembershipID = 0
	return nil
}
func DBUpdateCustomer(params *CustomerDB) error {
		if params.MembershipName == ""{
		if err := conn.Omit("membership_id").Updates(&params).Error; err != nil{
			return err	
		}
		return nil
	}
	var member_strt membership.MembershipVal 
	member_strt.Name= params.MembershipName
	membershipID,err := member_strt.GetID()
	if err !=nil{
		return err
	}
	params.MembershipID = membershipID
	status := conn.Updates(&params)
	params.MembershipID = 0
	if err := status.Error; err!=nil{
		return err
	}
	if status.RowsAffected == 0{
		return errors.New("no data with the input id")
	}
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
	params.MembershipID = 0
	return nil
}

func DBUpdateMembershipStatus(params *CustomerDB) error{
	var member_strt membership.MembershipVal 
	member_strt.Name= params.MembershipName
	membershipID,err := member_strt.GetID()
	if err !=nil{
		return err
	}
	params.MembershipID = membershipID
	status := conn.Model(&params).Update("membership_id",params.MembershipID)
	if err:= status.Error;err!=nil{
		return err
	}	
	if status.RowsAffected == 0{
		return errors.New("no data with the input id")
	}
	return nil
}
