package membership

import (
	"car-rental/utilities/db"
	//"fmt"
	//"log"

	"gorm.io/gorm"
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

func DBGetMembershipOne(params MembershipVal) MembershipVal{
	var result MembershipVal
	conn.First(&result,params.MembershipId)
	return result
}

func DBInsertMembership(params *MembershipVal) {
	err := conn.Create(&params).Error
	if err!=nil{
		panic(err)
	}
}
func DBUpdateMembership(params *MembershipVal) {
	err := conn.Updates(&params).Error
	if err!=nil{
		panic(err)
	}
}

func DBDeleteMembership(params *MembershipVal) {
	err := conn.Delete(&params).Error
	if err!=nil{
		panic(err)
	}
}
