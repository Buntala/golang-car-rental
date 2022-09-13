package service

import (
	"car-rental/entity"
	"car-rental/repository"
	"car-rental/request"
	"errors"
	"fmt"
)
type BookingService interface {
	FindAll()([]request.Booking)
	FindOne(booking request.Booking) (request.Booking,error)
	Save(booking request.Booking) (request.Booking,error)
	Update(booking request.Booking) (request.Booking,error)
	Delete(booking request.Booking) (request.Booking,error)
	SaveExtend(booking request.Booking) (request.Booking,error)
	SaveFinished(booking request.Booking) (request.Booking,error)
}
type bookingService struct{
	repository repository.BookingRepository
}

func NewBookingService(bookingRepository repository.BookingRepository) BookingService{
	return &bookingService{
		repository : bookingRepository,
	}
}

func (service *bookingService) FindAll() ([]request.Booking){
	dbBook := service.repository.FindAll()
	var res []request.Booking
	for i := range dbBook{
		row := request.DBtoReqBooking(dbBook[i])
		row.BookingTypeName = service.repository.GetBookingTypeName(dbBook[i])
		res = append(res, row)
	}
	return res
}

func (service *bookingService) FindOne(booking request.Booking) (request.Booking,error){
	if err := booking.Validate("get"); err!=nil{
		return booking,err
	}
	b_entity := booking.ToDB()
	err := service.repository.FindOne(&b_entity)
	res := request.DBtoReqBooking(b_entity)
	res.BookingTypeName = service.repository.GetBookingTypeName(b_entity)
	return res ,err
}

//
func (service *bookingService) Save(booking request.Booking) (request.Booking,error){
	if err :=  booking.Validate("post"); err!=nil{
		return booking,err
	}
	b_entity := booking.ToDB()
	var bookingType entity.BookingType
	bookingType.BookingType = booking.BookingTypeName
	b_entity.BookingTypeID = service.repository.GetBookingTypeID(bookingType)
	if err:= service.availabilityCheck(b_entity);err!=nil {
		return booking,err
	}
	service.calculate(&b_entity)
	err:=service.repository.Save(&b_entity)
	res := request.DBtoReqBooking(b_entity)
	res.BookingTypeName = booking.BookingTypeName
	return res ,err
}
//
func (service *bookingService) Update(booking request.Booking) (request.Booking,error){
	if err := booking.Validate("update"); err!=nil{
		return booking,err
	}
	b_entity := booking.ToDB()
	var bookingType entity.BookingType
	bookingType.BookingType = booking.BookingTypeName
	b_entity.BookingTypeID = service.repository.GetBookingTypeID(bookingType)

	var prev_data entity.Booking
	prev_data.BookingID = b_entity.BookingID
	if err := service.repository.FindOne(&prev_data);err!=nil{
		return booking,errors.New("no data with the input id")
	}
	if prev_data.Finished{
		return booking, errors.New("cannot update finished booking")
	}
	b_entity.CustomerID = prev_data.CustomerID
	b_entity.StartTime = prev_data.StartTime
	b_entity.EndTime = prev_data.EndTime
	if err:=service.availabilityCheck(b_entity);err!=nil{
		return booking,err
	}
	service.calculate(&b_entity)
	err:=service.repository.Update(&b_entity)
	res := request.DBtoReqBooking(b_entity)
	res.BookingTypeName = booking.BookingTypeName
	return res ,err
}

//
func (service *bookingService) Delete(booking request.Booking) (request.Booking,error){
	if err := booking.Validate("delete"); err!=nil{
		return booking,err
	}
	b_entity := booking.ToDB()
	err := service.repository.Delete(&b_entity)
	res := request.DBtoReqBooking(b_entity)
	res.BookingTypeName = service.repository.GetBookingTypeName(b_entity)
	return res ,err
}

func (service *bookingService) SaveExtend(booking request.Booking) (request.Booking,error){
	if err := booking.Validate("extend"); err!=nil{
		return booking,err
	}
	b_entity := booking.ToDB()
	var prev_data entity.Booking
	prev_data.BookingID = b_entity.BookingID
	if err := service.repository.FindOne(&prev_data);err!=nil{
		return booking,errors.New("no data with the input id")
	}
	if prev_data.Finished{
		return booking, errors.New("cannot extend finished booking")
	}
	if prev_data.EndTime.After(b_entity.EndTime){
		err_str := fmt.Sprintf("please insert data higher than %v",prev_data.EndTime.Format("2006-01-02"))
		return booking,errors.New(err_str)
	}
	err := service.repository.SaveExtend(&b_entity)
	res := request.DBtoReqBooking(b_entity)
	return res ,err
}
func (service *bookingService) SaveFinished(booking request.Booking) (request.Booking,error){
	if err := booking.Validate("finish"); err!=nil{
		return booking,err
	}
	b_entity := booking.ToDB()
	err := service.repository.SaveFinished(&b_entity)
	res := request.DBtoReqBooking(b_entity)
	return res ,err
}

func (service *bookingService) calculate(booking *entity.Booking){
	//PUT ON TO DB
	duration := int(booking.EndTime.Sub(booking.StartTime).Hours()/24) + 1
	var car entity.Car
	car.CarsID= booking.CarsID
	booking.TotalCost = duration * service.repository.GetCarCost(booking)
	var cust entity.Customer
	cust.CustomerID = booking.CustomerID
	booking.Discount = booking.TotalCost * service.repository.GetMembershipDiscount(booking) / 100
	if (booking.BookingTypeID == 2){
		var driver entity.Driver
		driver.DriverID = booking.DriverID
		driverCost := service.repository.GetDriverCost(booking)
		booking.TotalDriverCost = duration * driverCost
		booking.DriverIncentive = int(float64(booking.TotalCost) * 0.05)
	}
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
		if driverStatus > 0{
			return errors.New("driver is booked")
		}
		return err
	}
	return nil
}

