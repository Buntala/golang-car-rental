package driver

import (
	"car-rental/utilities/db"

	"gorm.io/gorm"
)
var(
	conn *gorm.DB = db.DbConnectGorm()
	//conn *sqlx.DB = db.DbConnect()
)
func DBGetDriverAll() []DriverVal{
	//conn.AutoMigrate(&DriverVal{})
	var result []DriverVal
	conn.Order("driver_id desc").Find(&result)
	return result
}

func DBGetDriverOne(params DriverVal) DriverVal{
	var result DriverVal
	conn.First(&result,params.DriverId)
	return result
}

func DBInsertDriver(params *DriverVal) {
	err := conn.Create(&params).Error
	if err!=nil{
		panic(err)
	}
}
func DBUpdateDriver(params *DriverVal) {
	err := conn.Updates(&params).Error
	if err!=nil{
		panic(err)
	}
}

func DBDeleteDriver(params *DriverVal) {
	err := conn.Delete(&params).Error
	if err!=nil{
		panic(err)
	}
}