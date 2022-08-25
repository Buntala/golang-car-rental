package service

import (
	"car-rental/entity"
	"car-rental/repository"
)
type CustomerService interface {
	FindAll()([]entity.Customer)
}
type customerService struct{
	repository repository.CustomerRepository
}

func NewCustomerService(customerRepository repository.CustomerRepository) CustomerService{
	return &customerService{
		repository : customerRepository,
	}
}

func (service *customerService) FindAll() ([]entity.Customer){
	return service.repository.FindAll()
}