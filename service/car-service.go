package service

import (
	"car-rental/repository"
	"car-rental/request"
)
type CarService interface {
	FindAll()([]request.Car)
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

func (service *carService) FindAll() ([]request.Car){
	dbCar := service.repository.FindAll()
	var cars []request.Car
	for i := range dbCar{
		row := request.DBtoReqCar(dbCar[i])
		cars = append(cars, row)
	}
	return cars
}

func (service *carService) FindOne(car request.Car) (request.Car,error){
	if err := car.Validate("get"); err!=nil{
		return car,err
	}
	c_entity := car.ToDB()
	err:=service.repository.FindOne(&c_entity)
	res := request.DBtoReqCar(c_entity)
	return res ,err
}

func (service *carService) Save(car request.Car) (request.Car,error){
	if err := car.Validate("post"); err!=nil{
		return car,err
	}
	c_entity := car.ToDB()
	err:=service.repository.Save(&c_entity)
	res := request.DBtoReqCar(c_entity)
	return res ,err
}
func (service *carService) Update(car request.Car) (request.Car,error){
	if err := car.Validate("update"); err!=nil{
		return car,err
	}
	c_entity := car.ToDB()
	err:=service.repository.Update(&c_entity)
	res := request.DBtoReqCar(c_entity)
	return res ,err
}

func (service *carService) Delete(car request.Car) (request.Car,error){
	if err := car.Validate("delete"); err!=nil{
		return car,err
	}
	c_entity := car.ToDB()
	err:=service.repository.Delete(&c_entity)
	res := request.DBtoReqCar(c_entity)
	return res ,err
}
