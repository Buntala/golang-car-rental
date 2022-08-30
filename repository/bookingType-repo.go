package repository

import (
	"car-rental/db"
	"car-rental/entity"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BookingTypeRepository interface {
	Save(bookingType entity.BookingType) (entity.BookingType,error)
	Update(bookingType entity.BookingType) (entity.BookingType,error)
	Delete(bookingType entity.BookingType) (entity.BookingType,error)
	FindAll() []entity.BookingType
	FindOne(bookingType entity.BookingType) (entity.BookingType,error)
}

type databaseBookingType struct {
	connection *gorm.DB
}

func NewBookingTypeRepository() BookingTypeRepository {
	return &databaseBookingType{
		connection: db.DB,
	}
}

func (db *databaseBookingType) Save(bookingType entity.BookingType)(entity.BookingType,error){
	status := db.connection.Clauses(clause.Returning{}).Create(&bookingType)
	if status.RowsAffected == 0{
		return bookingType, errors.New("no data with the id")
	}
	return bookingType,status.Error
}

func (db *databaseBookingType) Update(bookingType entity.BookingType) (entity.BookingType,error){
	status := db.connection.Updates(&bookingType)
	if status.RowsAffected == 0{
		return bookingType, errors.New("no data with the id")
	}
	return bookingType,status.Error
}

func (db *databaseBookingType) Delete(bookingType entity.BookingType) (entity.BookingType,error){
	status := db.connection.Delete(&bookingType)
	if status.RowsAffected == 0{
		return bookingType, errors.New("no data with the id")
	}
	return bookingType,status.Error
}

func (db *databaseBookingType) FindAll() []entity.BookingType {
	var bookingTypes []entity.BookingType
	db.connection.Set("gorm:auto_preload", true).Order("booking_type_id desc").Find(&bookingTypes)
	return bookingTypes
}

func (db *databaseBookingType) FindOne(bookingType entity.BookingType) (entity.BookingType,error){
	status := db.connection.Find(&bookingType,bookingType.BookingTypeID)
	if status.RowsAffected == 0{
		return bookingType, errors.New("no data with the id")
	}
	return bookingType, status.Error
}
/*
func (db *databaseBookingType) GetID(bookingType entity.BookingType) int{
	var result entity.BookingType
	status := db.connection.Where("name = ?", bookingType.Name).First(&result)
	return result.BookingTypeID
}*/
