package bookingType

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
func DBGetBookingTypeAll() []BookingTypeDB{
	var result []BookingTypeDB
	conn.Order("booking_type_id desc").Find(&result)
	return result
}

func DBGetBookingTypeOne(params BookingTypeDB) (BookingTypeDB,error){
	var result BookingTypeDB
	err := conn.First(&result,params.BookingTypeID).Error
	return result,err
}

func DBInsertBookingType(params *BookingTypeDB) error{
	err := conn.Create(&params).Error
	return err
}
func DBUpdateBookingType(params *BookingTypeDB) error{
	status := conn.Updates(&params)
	if err:= status.Error;err!=nil{
		return err
	}
	if status.RowsAffected == 0{
		return errors.New("no data with the input id")
	}
	return nil
}

func DBDeleteBookingType(params *BookingTypeDB) error{
	status := conn.Clauses(clause.Returning{}).Delete(&params)
	if err:= status.Error;err!=nil{
		return err
	}
	if status.RowsAffected == 0{
		return errors.New("no data with the input id")
	}
	return nil
}