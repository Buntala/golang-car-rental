package driver

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
func DBGetDriverAll() []DriverVal{
	//conn.AutoMigrate(&DriverVal{})
	var result []DriverVal
	conn.Order("driver_id desc").Find(&result)
	return result
}

func DBGetDriverOne(params DriverVal) (DriverVal,error){
	var result DriverVal
	err := conn.First(&result,params.DriverId).Error
	return result,err
}

func DBInsertDriver(params *DriverVal) error{
	err := conn.Create(&params).Error
	return err

}
func DBUpdateDriver(params *DriverVal) error{
	status := conn.Updates(&params)
	if err:= status.Error;err!=nil{
		return err
	}
	if status.RowsAffected == 0{
		return errors.New("no data with the input id")
	}
	return nil

}

func DBDeleteDriver(params *DriverVal) error{
	status := conn.Clauses(clause.Returning{}).Delete(&params)
	if err:= status.Error;err!=nil{
		return err
	}
	if status.RowsAffected == 0{
		return errors.New("no data with the input id")
	}
	return nil
}