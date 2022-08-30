package service

import (
	"car-rental/entity"
	"car-rental/repository"
	"car-rental/request"
)
type DriverService interface {
	FindAll()([]entity.Driver)
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

func (service *driverService) FindAll() ([]entity.Driver){
	return service.repository.FindAll()
}

func (service *driverService) FindOne(driver request.Driver) (request.Driver,error){
	driver.Validate("get")
	m_entity := driver.ToDB()
	res_entity,err:=service.repository.FindOne(m_entity)
	res := request.DBtoReqDriver(res_entity)
	return res ,err
}

func (service *driverService) Save(driver request.Driver) (request.Driver,error){
	driver.Validate("post")
	m_entity := driver.ToDB()
	res_entity,err:=service.repository.Save(m_entity)
	res := request.DBtoReqDriver(res_entity)
	return res ,err
}
func (service *driverService) Update(driver request.Driver) (request.Driver,error){
	driver.Validate("update")
	m_entity := driver.ToDB()
	res_entity,err:=service.repository.Update(m_entity)
	res := request.DBtoReqDriver(res_entity)
	return res ,err
}

func (service *driverService) Delete(driver request.Driver) (request.Driver,error){
	driver.Validate("delete")
	m_entity := driver.ToDB()
	res_entity,err:=service.repository.Delete(m_entity)
	res := request.DBtoReqDriver(res_entity)
	return res ,err
}
