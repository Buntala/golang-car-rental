package booking

import (
	"car-rental/utilities/db"
	"errors"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)
var(
	conn *gorm.DB = db.DbConnectGorm()
)
func DBGetBookingAll() []BookingDB{
	//conn.AutoMigrate(&BookingDB{})
	var result []BookingDB
	conn.Order("booking_id desc").Find(&result)
	return result
}

func DBGetBookingOne(params *BookingDB) (BookingDB,error){
	var result BookingDB
	err := conn.First(&result,params.BookingID).Error
	return result,err
}

func DBInsertBooking(params *BookingDB) error{
	if err:=params.availabilityCheck();err!=nil{
		return err
	}
	if err := params.calculate();err!=nil{
		return err
	}
	if params.BookingTypeName == "Car & Driver"{
		err := conn.Create(&params).Error
		return err
	}
	err := conn.Omit("driver_id,driver_incentive,total_driver_cost").Create(&params).Error
	return err

}
func DBUpdateBooking(params *BookingDB) error{
	if err:=params.availabilityCheck();err!=nil{
		return err
	}
	var result BookingDB
	result,err := DBGetBookingOne(params)
	if result.Finished{
		return errors.New("cannot update finished booking")
	}
	if err!=nil{
		return err
	}
	params.CustomerID = result.CustomerID
	if err := params.calculate();err!=nil{
		return err
	}
	if params.BookingTypeName == "Car & Driver"{
		err := conn.Omit("customer_id,start_time,end_time").Updates(&params).Error
		return err
	}
	status := conn.Omit("customer_id,start_time,end_time,driver_id,driver_incentive,total_driver_cost").Updates(&params)
	if err:= status.Error;err!=nil{
		return err
	}
	if status.RowsAffected == 0{
		return errors.New("no data with the input id")
	}
	return nil
}

func DBDeleteBooking(params *BookingDB) error{
	var finishCheck BookingDB
	conn.Take(&finishCheck,params.BookingID)
	if finishCheck.Finished{
		return errors.New("cannot delete finished booking")
	}
	status := conn.Clauses(clause.Returning{}).Delete(&params)
	if err:= status.Error;err!=nil{
		return err
	}
	if status.RowsAffected == 0{
		return errors.New("no data with the input id")
	}
	return nil
}

func DBFinished(params *BookingDB) error{
	status := conn.Clauses(clause.Returning{}).Model(&params).Update("finished", "true")
	if err:= status.Error;err!=nil{
		return err
	}
	if status.RowsAffected == 0{
		return errors.New("no data with the input id")
	}
	return nil
}


func DBExtend(params *BookingDB) error{
	if err:=params.availabilityCheck();err!=nil{
		return err
	}
	var prev_data BookingDB
	conn.Order("booking_id desc").Find(&prev_data)
	if prev_data.Finished{
		return errors.New("cannot extend finished booking")
	}
	if prev_data.EndTime.After(params.EndTime){
		err_str := fmt.Sprintf("Please insert data higher than %v",prev_data.EndTime.Format("2006-01-02"))
		return errors.New(err_str)
	}
	err := conn.Clauses(clause.Returning{}).Model(&params).Update("end_time", params.EndTime).Error
	return err
}