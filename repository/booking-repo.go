package repository

import (
	"car-rental/db"
	"car-rental/entity"
	"errors"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BookingRepository interface {
	Save(booking *entity.Booking) (error)
	Update(booking *entity.Booking) (error)
	Delete(booking *entity.Booking) (error)
	FindAll() []entity.Booking
	FindOne(booking *entity.Booking) (error)
	SaveExtend(booking *entity.Booking) (error)
	SaveFinished(booking *entity.Booking) (error)

	GetBookingTypeID(bookingType entity.BookingType) int
	GetBookingTypeName(booking entity.Booking) string
	GetCarCost(booking *entity.Booking) int
	GetDriverCost(booking *entity.Booking) int
	GetCarStatus(booking entity.Booking) (int, int, error)
	GetDriverStatus(booking entity.Booking) (int,error)
	GetMembershipDiscount(booking *entity.Booking) int
}

type databaseBooking struct {
	connection *gorm.DB
}

func NewBookingRepository() BookingRepository {
	return &databaseBooking{
		connection: db.DB,
	}
}

func (db *databaseBooking) Save(booking *entity.Booking)(error){
	status := db.connection.Clauses(clause.Returning{}).Create(&booking)
	if status.RowsAffected == 0{
		return errors.New("no data with the id")
	}
	return status.Error
}

func (db *databaseBooking) Update(booking *entity.Booking) (error){
	status := db.connection.Clauses(clause.Returning{}).Omit("customer_id,start_time,end_time").Updates(&booking)
	if status.RowsAffected == 0{
		return errors.New("no data with the id/can't update finished data")
	}
	return status.Error
}

func (db *databaseBooking) Delete(booking *entity.Booking) (error){
	status := db.connection.Clauses(clause.Returning{}).Delete(&booking)
	if status.RowsAffected == 0{
		return errors.New("no data with the id/can't update finished data")
	}
	return status.Error
}

func (db *databaseBooking) SaveFinished(booking *entity.Booking) (error){
	status := db.connection.Clauses(clause.Returning{}).Model(&booking).Update("finished", "true")
	if err:= status.Error;err!=nil{
		return err
	}
	if status.RowsAffected == 0{
		return errors.New("no data with the input id")
	}
	return nil
}

func (db *databaseBooking) SaveExtend(booking *entity.Booking) (error){
	var prev_data entity.Booking
	if err := db.connection.First(&prev_data,booking.BookingID).Error;err!=nil{
		return errors.New("no data with the input id")
	}
	if prev_data.Finished{
		return errors.New("cannot extend finished booking")
	}
	if prev_data.EndTime.After(booking.EndTime){
		err_str := fmt.Sprintf("please insert data higher than %v",prev_data.EndTime.Format("2006-01-02"))
		return errors.New(err_str)
	}
	status := db.connection.Clauses(clause.Returning{}).Model(&booking).Update("end_time", booking.EndTime)
	if err:= status.Error;err!=nil{	
		return err
	}
	if status.RowsAffected == 0{
		return errors.New("no data with the input id")
	}
	return nil
}

func (db *databaseBooking) FindAll() []entity.Booking {
	var bookings []entity.Booking
	db.connection.Set("gorm:auto_preload", true).Order("booking_id desc").Find(&bookings)
	return bookings
}

func (db *databaseBooking) FindOne(booking *entity.Booking) (error){
	status := db.connection.First(&booking,booking.BookingID)
	if status.RowsAffected == 0{
		return errors.New("no data with the id")
	}
	return status.Error
}


// non api - functions
func (db *databaseBooking) GetBookingTypeID(bookingType entity.BookingType) int{
	//var result entity.BookingType
	db.connection.Where("booking_type = ?", bookingType.BookingType).First(&bookingType)
	return bookingType.BookingTypeID
}
func (db *databaseBooking) GetBookingTypeName(booking entity.Booking) string{
	var bookingType entity.BookingType
	db.connection.First(&bookingType,booking.BookingTypeID)
	return bookingType.BookingType
}

func (db *databaseBooking) GetCarCost(booking *entity.Booking) int{
	var result entity.Car
	db.connection.First(&result,booking.CarsID)
	return result.RentPriceDaily
}

func (db *databaseBooking) GetDriverCost(booking *entity.Booking) int{
	var result entity.Driver
	db.connection.First(&result,booking.DriverID)
	return result.DailyCost
}

func (db *databaseBooking) GetCarStatus(booking entity.Booking) (int, int, error){
	var car entity.Car
	success := db.connection.Find(&car,booking.CarsID).RowsAffected
	if success == 0{
		return 0,0, errors.New("car id is invalid")
	}
	var carBooked []entity.Booking
	booked:= db.connection.Where("cars_id= ?",booking.CarsID).Where("booking_id != ?",booking.BookingID).Where(
	db.connection.Where(db.connection.Where(
	"start_time >= ?",booking.StartTime).Where("start_time <= ?",booking.EndTime)).Or(db.connection.Where(
	"end_time >= ?",booking.StartTime).Where("end_time <= ?",booking.EndTime)).Or(db.connection.Where(
	"start_time <= ?" , booking.StartTime).Where("end_time >= ?" , booking.EndTime))).Find(&carBooked).RowsAffected
	return int(booked),car.Stock,nil
}

func (db *databaseBooking) GetDriverStatus(booking entity.Booking) (int,error){
	var driver entity.Driver
	success := db.connection.Find(&driver,booking.DriverID).RowsAffected
	if success == 0{
		return 0, errors.New("driver id is invalid")
	}
	var driverBook entity.Booking
	booked:= db.connection.Where("driver_id= ?",booking.DriverID).Where(db.connection.Where(db.connection.Where(
		"start_time >= ?",booking.StartTime).Where("start_time <= ?",booking.EndTime)).Or(db.connection.Where(
		"end_time >= ?",booking.StartTime).Where("end_time <= ?",booking.EndTime)).Or(db.connection.Where(
		"start_time <= ?" , booking.StartTime).Where("end_time >= ?" , booking.EndTime))).Find(&driverBook).RowsAffected
	return int(booked),nil
}

func (db *databaseBooking) GetMembershipDiscount(booking *entity.Booking) int{
	var customer entity.Customer
	var membership entity.Membership
	db.connection.First(&customer,booking.CustomerID)
	if customer.MembershipID ==0 {
		return 0
	}
	db.connection.First(&membership,customer.MembershipID)
	return membership.Discount
}
