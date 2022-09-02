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
	if err := customer.Validate("get"); err!=nil{
		return customer ,err
	}
	ct_entity := customer.ToDB()
	err  := service.repository.FindOne(&ct_entity)
	var membership entity.Membership

	cust_entity := request.DBtoReqCust(ct_entity)
	if ct_entity.MembershipID!=0{
		membership.MembershipID = ct_entity.MembershipID
		cust_entity.MembershipName = service.repository.GetMembershipName(membership)
	}
	return cust_entity,err
}

func (service *customerService) Save(customer request.CustomerRequest) (request.CustomerRequest,error){
	if err := customer.Validate("post"); err!=nil{
		return customer ,err
	}
	var mem_entity entity.Membership
	mem_entity.Name = customer.MembershipName
	ct_entity := customer.ToDB()
	if customer.MembershipName != ""{
		ct_entity.MembershipID = service.repository.GetMembershipID(mem_entity)
	}
	err  := service.repository.Save(&ct_entity)
	res:= request.DBtoReqCust(ct_entity)
	res.MembershipName = customer.MembershipName
	return res,err
}

func (service *customerService) Update(customer request.CustomerRequest) (request.CustomerRequest,error){
	if err := customer.Validate("update"); err!=nil{
		return customer ,err
	}
	var mem_entity entity.Membership
	mem_entity.Name = customer.MembershipName

	ct_entity := customer.ToDB()
	if customer.MembershipName != ""{
		ct_entity.MembershipID = service.repository.GetMembershipID(mem_entity)
	}
	err := service.repository.Update(&ct_entity)
	res:= request.DBtoReqCust(ct_entity)
	
	res.MembershipName = customer.MembershipName
	return res,err
}

func (service *customerService) Delete(customer request.CustomerRequest) (request.CustomerRequest,error){
	if err := customer.Validate("delete"); err!=nil{
		return customer ,err
	}
	ct_entity := customer.ToDB()
	err  := service.repository.Delete(&ct_entity)
	res:= request.DBtoReqCust(ct_entity)
	if ct_entity.MembershipID!=0{
		var membership entity.Membership
		membership.MembershipID = ct_entity.MembershipID
		res.MembershipName = service.repository.GetMembershipName(membership)
	}
	//res.MembershipName = customer.MembershipName
	return res,err
}

func (service *customerService) SaveMembership(customer request.CustomerRequest) (request.CustomerRequest,error){
	if err := customer.Validate("membership"); err!=nil{
		return customer ,err
	}
	ct_entity := customer.ToDB()
	err := service.repository.SaveMembership(&ct_entity)
	res:= request.DBtoReqCust(ct_entity)
	res.MembershipName = customer.MembershipName
	return res , err
}