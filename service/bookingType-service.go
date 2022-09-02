package service

import (
	"car-rental/repository"
	"car-rental/request"
)
type BookingTypeService interface {
	FindAll()([]request.BookingType)
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

func (service *bookingTypeService) FindAll() ([]request.BookingType){
	dbBookingType := service.repository.FindAll()
	var bookingTypes []request.BookingType
	for i := range dbBookingType{
		row := request.DBtoReqBookingType(dbBookingType[i])
		bookingTypes = append(bookingTypes, row)
	}
	return bookingTypes
}

func (service *bookingTypeService) FindOne(bookingType request.BookingType) (request.BookingType,error){
	if err := bookingType.Validate("get"); err!=nil{
		return bookingType,err
	}
	bt_entity := bookingType.ToDB()
	err:=service.repository.FindOne(&bt_entity)
	res := request.DBtoReqBookingType(bt_entity)
	return res ,err
}

func (service *bookingTypeService) Save(bookingType request.BookingType) (request.BookingType,error){
	if err := bookingType.Validate("post"); err!=nil{
		return bookingType,err
	}
	bt_entity := bookingType.ToDB()
	err:=service.repository.Save(&bt_entity)
	res := request.DBtoReqBookingType(bt_entity)
	return res ,err
}
func (service *bookingTypeService) Update(bookingType request.BookingType) (request.BookingType,error){
	if err := bookingType.Validate("update"); err!=nil{
		return bookingType,err
	}
	bt_entity := bookingType.ToDB()
	err:=service.repository.Update(&bt_entity)
	res := request.DBtoReqBookingType(bt_entity)
	return res ,err
}

func (service *bookingTypeService) Delete(bookingType request.BookingType) (request.BookingType,error){
	if err := bookingType.Validate("delete"); err!=nil{
		return bookingType,err
	}
	bt_entity := bookingType.ToDB()
	err:=service.repository.Delete(&bt_entity)
	res := request.DBtoReqBookingType(bt_entity)
	return res ,err
}
