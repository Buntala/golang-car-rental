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
	Save(booking entity.Booking) (entity.Booking,error)
	Update(booking entity.Booking) (entity.Booking,error)
	Delete(booking entity.Booking) (entity.Booking,error)
	FindAll() []entity.Booking
	FindOne(booking entity.Booking) (entity.Booking,error)
	SaveExtend(booking entity.Booking) (entity.Booking,error)

	GetBookingTypeID(bookingType entity.BookingType) int
	GetCarCost(booking entity.Booking) int
	GetDriverCost(booking entity.Booking) int
	GetCarStatus(booking entity.Booking) (int, int, error)
	GetDriverStatus(booking entity.Booking) (int,error)
	GetMembershipDiscount(booking entity.Booking) int
}

type databaseBooking struct {
	connection *gorm.DB
}

func NewBookingRepository() BookingRepository {
	return &databaseBooking{
		connection: db.DB,
	}
}

func (db *databaseBooking) Save(booking entity.Booking)(entity.Booking,error){
	status := db.connection.Clauses(clause.Returning{}).Create(&booking)
	if status.RowsAffected == 0{
		return booking, errors.New("no data with the id")
	}
	return booking,status.Error
}

func (db *databaseBooking) Update(booking entity.Booking) (entity.Booking,error){
	status := db.connection.Clauses(clause.Returning{}).Updates(&booking)
	if status.RowsAffected == 0{
		return booking, errors.New("no data with the id/can't update finished data")
	}
	return booking,status.Error
}

func (db *databaseBooking) Delete(booking entity.Booking) (entity.Booking,error){
	status := db.connection.Clauses(clause.Returning{}).Delete(&booking)
	if status.RowsAffected == 0{
		return booking, errors.New("no data with the id/can't update finished data")
	}
	return booking,status.Error
}

func (db *databaseBooking) SaveFinished(booking entity.Booking) (entity.Booking,error){
	status := db.connection.Clauses(clause.Returning{}).Model(&booking).Update("finished", "true")
	if err:= status.Error;err!=nil{
		return booking, err
	}
	if status.RowsAffected == 0{
		return booking, errors.New("no data with the input id")
	}
	return booking, nil
}
func (db *databaseBooking) SaveExtend(booking entity.Booking) (entity.Booking,error){
	/*
	if err:=params.availabilityCheck();err!=nil{
		return err
	}*/
	var prev_data entity.Booking
	db.connection.Find(&prev_data,booking.BookingID)
	if prev_data.Finished{
		return booking, errors.New("cannot extend finished booking")
	}
	if prev_data.EndTime.After(booking.EndTime){
		err_str := fmt.Sprintf("please insert data higher than %v",prev_data.EndTime.Format("2006-01-02"))
		return booking, errors.New(err_str)
	}
	err := db.connection.Clauses(clause.Returning{}).Model(&booking).Update("end_time", booking.EndTime).Error
	return booking, err
}

func (db *databaseBooking) FindAll() []entity.Booking {
	var bookings []entity.Booking
	db.connection.Set("gorm:auto_preload", true).Order("booking_id desc").Find(&bookings)
	return bookings
}

func (db *databaseBooking) FindOne(booking entity.Booking) (entity.Booking,error){
	status := db.connection.Find(&booking,booking.BookingID)
	if status.RowsAffected == 0{
		return booking, errors.New("no data with the id")
	}
	return booking, status.Error
}


// non api - functions
func (db *databaseBooking) GetBookingTypeID(bookingType entity.BookingType) int{
	//var result entity.BookingType
	db.connection.Where("name = ?", bookingType.BookingType).First(&bookingType)
	return bookingType.BookingTypeID
}

func (db *databaseBooking) GetCarCost(booking entity.Booking) int{
	var result entity.Car
	db.connection.First(&result,booking.CarsID)
	return result.RentPriceDaily
}

func (db *databaseBooking) GetDriverCost(booking entity.Booking) int{
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

func (db *databaseBooking) GetMembershipDiscount(booking entity.Booking) int{
	var customer entity.Customer
	var membership entity.Membership
	db.connection.First(&customer,booking.DriverID)
	if customer.MembershipID ==0 {
		return 0
	}
	db.connection.First(&membership,customer.MembershipID)
	return membership.Discount
}