package repository

import (
	"car-rental/db"
	"car-rental/entity"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DriverRepository interface {
	Save(driver entity.Driver) (entity.Driver,error)
	Update(driver entity.Driver) (entity.Driver,error)
	Delete(driver entity.Driver) (entity.Driver,error)
	FindAll() []entity.Driver
	FindOne(driver entity.Driver) (entity.Driver,error)
}

type databaseDriver struct {
	connection *gorm.DB
}

func NewDriverRepository() DriverRepository {
	return &databaseDriver{
		connection: db.DB,
	}
}

func (db *databaseDriver) Save(driver entity.Driver)(entity.Driver,error){
	status := db.connection.Clauses(clause.Returning{}).Create(&driver)
	if status.RowsAffected == 0{
		return driver, errors.New("no data with the id")
	}
	return driver,status.Error
}

func (db *databaseDriver) Update(driver entity.Driver) (entity.Driver,error){
	status := db.connection.Updates(&driver)
	if status.RowsAffected == 0{
		return driver, errors.New("no data with the id")
	}
	return driver,status.Error
}

func (db *databaseDriver) Delete(driver entity.Driver) (entity.Driver,error){
	status := db.connection.Delete(&driver)
	if status.RowsAffected == 0{
		return driver, errors.New("no data with the id")
	}
	return driver,status.Error
}

func (db *databaseDriver) FindAll() []entity.Driver {
	var drivers []entity.Driver
	db.connection.Set("gorm:auto_preload", true).Order("driver_id desc").Find(&drivers)
	return drivers
}

func (db *databaseDriver) FindOne(driver entity.Driver) (entity.Driver,error){
	status := db.connection.Find(&driver,driver.DriverID)
	if status.RowsAffected == 0{
		return driver, errors.New("no data with the id")
	}
	return driver, status.Error
}
/*
func (db *databaseDriver) GetID(driver entity.Driver) int{
	var result entity.Driver
	status := db.connection.Where("name = ?", driver.Name).First(&result)
	return result.DriverID
}*/
