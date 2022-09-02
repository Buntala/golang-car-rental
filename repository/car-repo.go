package repository

import (
	"car-rental/db"
	"car-rental/entity"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CarRepository interface {
	Save(car *entity.Car) (error)
	Update(car *entity.Car) (error)
	Delete(car *entity.Car) (error)
	FindAll() []entity.Car
	FindOne(car *entity.Car) (error)
}

type databaseCar struct {
	connection *gorm.DB
}

func NewCarRepository() CarRepository {
	return &databaseCar{
		connection: db.DB,
	}
}

func (db *databaseCar) Save(car *entity.Car)(error){
	status := db.connection.Clauses(clause.Returning{}).Create(&car)
	if status.RowsAffected == 0{
		return errors.New("no data with the id")
	}
	return status.Error
}

func (db *databaseCar) Update(car *entity.Car) (error){
	status := db.connection.Clauses(clause.Returning{}).Updates(&car)
	if status.Error!=nil{
		return status.Error
	}
	if status.RowsAffected == 0{
		return errors.New("no data with the id")
	}
	return status.Error
}

func (db *databaseCar) Delete(car *entity.Car) (error){
	status := db.connection.Clauses(clause.Returning{}).Delete(&car)
	if status.Error!=nil{
		return status.Error
	}
	if status.RowsAffected == 0{
		return errors.New("no data with the id")
	}
	return status.Error
}

func (db *databaseCar) FindAll() []entity.Car {
	var cars []entity.Car
	db.connection.Set("gorm:auto_preload", true).Order("cars_id desc").Find(&cars)
	return cars
}

func (db *databaseCar) FindOne(car *entity.Car) (error){
	status := db.connection.Find(&car,car.CarsID)
	if status.RowsAffected == 0{
		return errors.New("no data with the id")
	}
	return status.Error
}