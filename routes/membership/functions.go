package membership

import (
	"car-rental/utilities/db"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)
var(
	conn *gorm.DB = db.DbConnectGorm()
	//conn *sqlx.DB = db.DbConnect()
)
func DBGetMembershipAll() []MembershipVal{
	conn.AutoMigrate(&MembershipVal{})
	var result []MembershipVal
	conn.Order("membership_id desc").Find(&result)
	return result
}

func DBGetMembershipOne(params MembershipVal) (MembershipVal,error){
	var result MembershipVal
	err:= conn.First(&result,params.MembershipID).Error
	return result, err
}

func DBInsertMembership(params *MembershipVal) error{
	err := conn.Create(&params).Error
	return err
}
func DBUpdateMembership(params *MembershipVal) error{
	status := conn.Updates(&params)
	if err:= status.Error;err!=nil{
		return err
	}
	if status.RowsAffected == 0{
		return errors.New("no data with the input id")
	}
	return nil
}

func DBDeleteMembership(params *MembershipVal) error{
	status := conn.Clauses(clause.Returning{}).Delete(&params)
	if err:= status.Error;err!=nil{
		return err
	}
	if status.RowsAffected == 0{
		return errors.New("no data with the input id")
	}
	return nil
}
