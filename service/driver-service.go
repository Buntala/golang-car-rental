package service

import (
	"car-rental/repository"
	"car-rental/request"
)
type DriverService interface {
	FindAll()([]request.Driver)
	FindOne(driver request.Driver) (request.Driver,error)
	Save(driver request.Driver) (request.Driver,error)
	Update(driver request.Driver) (request.Driver,error)
	Delete(driver request.Driver) (request.Driver,error)
}
type driverService struct{
	repository repository.DriverRepository
}

func NewDriverService(driverRepository repository.DriverRepository) DriverService{
	return &driverService{
		repository : driverRepository,
	}
}

func (service *driverService) FindAll() ([]request.Driver){
	dbDriver := service.repository.FindAll()
	var drivers []request.Driver
	for i := range dbDriver{
		row := request.DBtoReqDriver(dbDriver[i])
		drivers = append(drivers, row)
	}
	return drivers
}

func (service *driverService) FindOne(driver request.Driver) (request.Driver,error){
	if err := driver.Validate("get"); err!=nil{
		return driver ,err
	}
	d_entity := driver.ToDB()
	err:=service.repository.FindOne(&d_entity)
	res := request.DBtoReqDriver(d_entity)
	return res ,err
}

func (service *driverService) Save(driver request.Driver) (request.Driver,error){
	if err := driver.Validate("post"); err!=nil{
		return driver ,err
	}
	d_entity := driver.ToDB()
	err:=service.repository.Save(&d_entity)
	res := request.DBtoReqDriver(d_entity)
	return res ,err
}
func (service *driverService) Update(driver request.Driver) (request.Driver,error){
	if err := driver.Validate("update"); err!=nil{
		return driver ,err
	}
	d_entity := driver.ToDB()
	err:=service.repository.Update(&d_entity)
	res := request.DBtoReqDriver(d_entity)
	return res ,err
}

func (service *driverService) Delete(driver request.Driver) (request.Driver,error){
	if err := driver.Validate("delete"); err!=nil{
		return driver ,err
	}
	d_entity := driver.ToDB()
	err:=service.repository.Delete(&d_entity)
	res := request.DBtoReqDriver(d_entity)
	return res ,err
}
