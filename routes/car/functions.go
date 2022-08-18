package car

import (
	"car-rental/utilities/db"

	"gorm.io/gorm"
)
var(
	conn *gorm.DB = db.DbConnectGorm()
	//conn *sqlx.DB = db.DbConnect()
)
func DBGetCarAll() []CarDB{
	//conn.AutoMigrate(&CarDB{})
	var result []CarDB
	conn.Order("driver_id desc").Find(&result)
	return result
}

func DBGetCarOne(params CarDB) CarDB{
	var result CarDB
	conn.First(&result,params.CarsID)
	return result
}

func DBInsertCar(params *CarDB) {
	err := conn.Create(&params).Error
	if err!=nil{
		panic(err)
	}
}
func DBUpdateCar(params *CarDB) {
	err := conn.Updates(&params).Error
	if err!=nil{
		panic(err)
	}
}

func DBDeleteCar(params *CarDB) {
	err := conn.Delete(&params).Error
	if err!=nil{
		panic(err)
	}
}