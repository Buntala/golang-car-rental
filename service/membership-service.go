package service

import (
	"car-rental/entity"
	"car-rental/repository"
	"car-rental/request"
)
type MembershipService interface {
	FindAll()([]entity.Membership)
	FindOne(membership request.Membership) (request.Membership,error)
	Save(membership request.Membership) (request.Membership,error)
	Update(membership request.Membership) (request.Membership,error)
	Delete(membership request.Membership) (request.Membership,error)
}
type membershipService struct{
	repository repository.MembershipRepository
}

func NewMembershipService(membershipRepository repository.MembershipRepository) MembershipService{
	return &membershipService{
		repository : membershipRepository,
	}
}

func (service *membershipService) FindAll() ([]entity.Membership){
	return service.repository.FindAll()
}

func (service *membershipService) FindOne(membership request.Membership) (request.Membership,error){
	membership.Validate("get")
	m_entity := membership.ToDB()
	res_entity,err:=service.repository.FindOne(m_entity)
	res := request.DBtoReqMember(res_entity)
	return res ,err
}

func (service *membershipService) Save(membership request.Membership) (request.Membership,error){
	membership.Validate("post")
	m_entity := membership.ToDB()
	res_entity,err:=service.repository.Save(m_entity)
	res := request.DBtoReqMember(res_entity)
	return res ,err
}
func (service *membershipService) Update(membership request.Membership) (request.Membership,error){
	membership.Validate("update")
	m_entity := membership.ToDB()
	res_entity,err:=service.repository.Update(m_entity)
	res := request.DBtoReqMember(res_entity)
	return res ,err
}

func (service *membershipService) Delete(membership request.Membership) (request.Membership,error){
	membership.Validate("delete")
	m_entity := membership.ToDB()
	res_entity,err:=service.repository.Delete(m_entity)
	res := request.DBtoReqMember(res_entity)
	return res ,err
}
