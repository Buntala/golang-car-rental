package service

import (
	"car-rental/entity"
	"car-rental/repository"
	"car-rental/request"
	"errors"
)
type BookingService interface {
	FindAll()([]entity.Booking)
	FindOne(booking request.Booking) (request.Booking,error)
	Save(booking request.Booking) (request.Booking,error)
	Update(booking request.Booking) (request.Booking,error)
	Delete(booking request.Booking) (request.Booking,error)
}
type bookingService struct{
	repository repository.BookingRepository
}

func NewBookingService(bookingRepository repository.BookingRepository) BookingService{
	return &bookingService{
		repository : bookingRepository,
	}
}

func (service *bookingService) FindAll() ([]entity.Booking){
	return service.repository.FindAll()
}

func (service *bookingService) FindOne(booking request.Booking) (request.Booking,error){
	booking.Validate("get")
	m_entity := booking.ToDB()
	res_entity,err:=service.repository.FindOne(m_entity)
	res := request.DBtoReqBooking(res_entity)
	return res ,err
}

//
func (service *bookingService) Save(booking request.Booking) (request.Booking,error){
	booking.Validate("post")
	m_entity := booking.ToDB()
	res_entity,err:=service.repository.Save(m_entity)
	res := request.DBtoReqBooking(res_entity)
	return res ,err
}
//
func (service *bookingService) Update(booking request.Booking) (request.Booking,error){
	booking.Validate("update")
	m_entity := booking.ToDB()
	service.calculate(&m_entity)
	res_entity,err:=service.repository.Update(m_entity)
	res := request.DBtoReqBooking(res_entity)
	return res ,err
}

//
func (service *bookingService) Delete(booking request.Booking) (request.Booking,error){
	booking.Validate("delete")
	m_entity := booking.ToDB()
	res_entity,err:=service.repository.Delete(m_entity)
	res := request.DBtoReqBooking(res_entity)
	return res ,err
}

func (service *bookingService) SaveExtend(booking request.Booking) (request.Booking,error){
	booking.Validate("extend")
	m_entity := booking.ToDB()
	var bookingType entity.BookingType
	bookingType.BookingType = booking.BookingTypeName
	m_entity.BookingTypeID = service.repository.GetBookingTypeID(bookingType)
	service.availabilityCheck(m_entity)
	service.calculate(&m_entity)
	res_entity,err:=service.repository.SaveExtend(m_entity)
	res := request.DBtoReqBooking(res_entity)
	return res ,err
}

func (service *bookingService) calculate(booking *entity.Booking){

}
func (service *bookingService) availabilityCheck(booking entity.Booking) error{
	booked,stock,err:= service.repository.GetCarStatus(booking)
	if err!=nil{
		return err
	}
	if stock <= booked{
		return errors.New("car is fully booked")
	}
	if booking.BookingTypeID == 2 {
		driverStatus,err := service.repository.GetDriverStatus(booking)
		if err!=nil{
			return err
		}
		if driverStatus <= 0{
			return errors.New("driver is booked")
		}
		return err
	}
	return nil
}

