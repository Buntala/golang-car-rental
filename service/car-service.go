package service

import (
	"car-rental/entity"
	"car-rental/repository"
	"car-rental/request"
)
type CarService interface {
	FindAll()([]entity.Car)
	FindOne(car request.Car) (request.Car,error)
	Save(car request.Car) (request.Car,error)
	Update(car request.Car) (request.Car,error)
	Delete(car request.Car) (request.Car,error)
}
type carService struct{
	repository repository.CarRepository
}

func NewCarService(carRepository repository.CarRepository) CarService{
	return &carService{
		repository : carRepository,
	}
}

func (service *carService) FindAll() ([]entity.Car){
	return service.repository.FindAll()
}

func (service *carService) FindOne(car request.Car) (request.Car,error){
	car.Validate("get")
	m_entity := car.ToDB()
	res_entity,err:=service.repository.FindOne(m_entity)
	res := request.DBtoReqCar(res_entity)
	return res ,err
}

func (service *carService) Save(car request.Car) (request.Car,error){
	car.Validate("post")
	m_entity := car.ToDB()
	res_entity,err:=service.repository.Save(m_entity)
	res := request.DBtoReqCar(res_entity)
	return res ,err
}
func (service *carService) Update(car request.Car) (request.Car,error){
	car.Validate("update")
	m_entity := car.ToDB()
	res_entity,err:=service.repository.Update(m_entity)
	res := request.DBtoReqCar(res_entity)
	return res ,err
}

func (service *carService) Delete(car request.Car) (request.Car,error){
	car.Validate("delete")
	m_entity := car.ToDB()
	res_entity,err:=service.repository.Delete(m_entity)
	res := request.DBtoReqCar(res_entity)
	return res ,err
}
