package service

import (
	"car-rental/entity"
	"car-rental/repository"
	"car-rental/request"
)
type CustomerService interface {
	FindAll()([]request.CustomerRequest)
	FindOne(customer request.CustomerRequest) (request.CustomerRequest,error)
	Save(customer request.CustomerRequest) (request.CustomerRequest,error)
	Update(customer request.CustomerRequest) (request.CustomerRequest,error)
	Delete(customer request.CustomerRequest) (request.CustomerRequest,error)
	SaveMembership(customer request.CustomerRequest) (request.CustomerRequest,error)
}
type customerService struct{
	repository repository.CustomerRepository
}

func NewCustomerService(customerRepository repository.CustomerRepository) CustomerService{
	return &customerService{
		repository : customerRepository,
	}
}

func (service *customerService) FindAll() ([]request.CustomerRequest){
	dbCust := service.repository.FindAll()
	var res []request.CustomerRequest
	var membership entity.Membership
	for i := range dbCust{
		row := request.DBtoReqCust(dbCust[i])
		if dbCust[i].MembershipID!=0{
			membership.MembershipID = dbCust[i].MembershipID
			row.MembershipName = service.repository.GetMembershipName(membership)
		}
		res = append(res, row)
	}
	return res
}
func (service *customerService) FindOne(customer request.CustomerRequest) (request.CustomerRequest,error){
	customer.Validate("get")
	ct_entity := customer.ToDB()
	dbCust,err  := service.repository.FindOne(ct_entity)
	var membership entity.Membership

	cust_entity := request.DBtoReqCust(dbCust)
	if dbCust.MembershipID!=0{
		membership.MembershipID = dbCust.MembershipID
		cust_entity.MembershipName = service.repository.GetMembershipName(membership)
	}
	return cust_entity,err
}

func (service *customerService) Save(customer request.CustomerRequest) (request.CustomerRequest,error){
	customer.Validate("post")
	var mem_entity entity.Membership
	mem_entity.Name = customer.MembershipName
	ct_entity := customer.ToDB()
	if customer.MembershipName != ""{
		ct_entity.MembershipID = service.repository.GetMembershipID(mem_entity)
	}
	dbCust,err  := service.repository.Save(ct_entity)
	res:= request.DBtoReqCust(dbCust)
	res.MembershipName = customer.MembershipName
	return res,err
}

func (service *customerService) Update(customer request.CustomerRequest) (request.CustomerRequest,error){
	customer.Validate("update")
	var mem_entity entity.Membership
	mem_entity.Name = customer.MembershipName

	ct_entity := customer.ToDB()
	if customer.MembershipName != ""{
		ct_entity.MembershipID = service.repository.GetMembershipID(mem_entity)
	}
	dbCust,err := service.repository.Update(ct_entity)
	res:= request.DBtoReqCust(dbCust)
	res.MembershipName = customer.MembershipName
	return res,err
}

func (service *customerService) Delete(customer request.CustomerRequest) (request.CustomerRequest,error){
	customer.Validate("delete")
	ct_entity := customer.ToDB()
	dbCust,err  := service.repository.Delete(ct_entity)
	res:= request.DBtoReqCust(dbCust)
	res.MembershipName = customer.MembershipName
	return res,err
}

func (service *customerService) SaveMembership(customer request.CustomerRequest) (request.CustomerRequest,error){
	err:=customer.Validate("membership")
	if err!=nil{
		return customer ,err
	}
	ct_entity := customer.ToDB()
	dbCust,err := service.repository.SaveMembership(ct_entity)
	res:= request.DBtoReqCust(dbCust)
	res.MembershipName = customer.MembershipName
	return res , err
}