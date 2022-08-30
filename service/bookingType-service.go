package service

import (
	"car-rental/entity"
	"car-rental/repository"
	"car-rental/request"
)
type BookingTypeService interface {
	FindAll()([]entity.BookingType)
	FindOne(bookingType request.BookingType) (request.BookingType,error)
	Save(bookingType request.BookingType) (request.BookingType,error)
	Update(bookingType request.BookingType) (request.BookingType,error)
	Delete(bookingType request.BookingType) (request.BookingType,error)
}
type bookingTypeService struct{
	repository repository.BookingTypeRepository
}

func NewBookingTypeService(bookingTypeRepository repository.BookingTypeRepository) BookingTypeService{
	return &bookingTypeService{
		repository : bookingTypeRepository,
	}
}

func (service *bookingTypeService) FindAll() ([]entity.BookingType){
	return service.repository.FindAll()
}

func (service *bookingTypeService) FindOne(bookingType request.BookingType) (request.BookingType,error){
	bookingType.Validate("get")
	m_entity := bookingType.ToDB()
	res_entity,err:=service.repository.FindOne(m_entity)
	res := request.DBtoReqBookingType(res_entity)
	return res ,err
}

func (service *bookingTypeService) Save(bookingType request.BookingType) (request.BookingType,error){
	bookingType.Validate("post")
	m_entity := bookingType.ToDB()
	res_entity,err:=service.repository.Save(m_entity)
	res := request.DBtoReqBookingType(res_entity)
	return res ,err
}
func (service *bookingTypeService) Update(bookingType request.BookingType) (request.BookingType,error){
	bookingType.Validate("update")
	m_entity := bookingType.ToDB()
	res_entity,err:=service.repository.Update(m_entity)
	res := request.DBtoReqBookingType(res_entity)
	return res ,err
}

func (service *bookingTypeService) Delete(bookingType request.BookingType) (request.BookingType,error){
	bookingType.Validate("delete")
	m_entity := bookingType.ToDB()
	res_entity,err:=service.repository.Delete(m_entity)
	res := request.DBtoReqBookingType(res_entity)
	return res ,err
}
