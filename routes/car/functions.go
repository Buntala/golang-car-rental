package car

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
func DBGetCarAll() []CarDB{
	//conn.AutoMigrate(&CarDB{})
	var result []CarDB
	conn.Order("cars_id desc").Find(&result)
	return result
}

func DBGetCarOne(params CarDB) (CarDB, error){
	var result CarDB
	err := conn.First(&result,params.CarsID).Error
	return result,err
}

func DBInsertCar(params *CarDB) error {
	err := conn.Create(&params).Error
	return err
}
func DBUpdateCar(params *CarDB) error{
	status := conn.Updates(&params)
	if err:= status.Error;err!=nil{
		return err
	}
	if status.RowsAffected == 0{
		return errors.New("no data with the input id")
	}
	return nil
}

func DBDeleteCar(params *CarDB) error{
	status := conn.Clauses(clause.Returning{}).Delete(&params)
	if err:= status.Error;err!=nil{
		return err
	}
	if status.RowsAffected == 0{
		return errors.New("no data with the input id")
	}
	return nil
}