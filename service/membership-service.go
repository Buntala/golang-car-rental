package service

import (
	"car-rental/repository"
	"car-rental/request"
)
type MembershipService interface {
	FindAll()([]request.Membership)
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

func (service *membershipService) FindAll() ([]request.Membership){
	dbMembership := service.repository.FindAll()
	var memberships []request.Membership
	for i := range dbMembership{
		row := request.DBtoReqMember(dbMembership[i])
		memberships = append(memberships, row)
	}
	return memberships
}

func (service *membershipService) FindOne(membership request.Membership) (request.Membership,error){
	if err := membership.Validate("get"); err!=nil{
		return membership ,err
	}
	m_entity := membership.ToDB()
	err:=service.repository.FindOne(&m_entity)
	res := request.DBtoReqMember(m_entity)
	return res ,err
}

func (service *membershipService) Save(membership request.Membership) (request.Membership,error){
	if err := membership.Validate("post"); err!=nil{
		return membership ,err
	}
	m_entity := membership.ToDB()
	err:=service.repository.Save(&m_entity)
	res := request.DBtoReqMember(m_entity)
	return res ,err
}
func (service *membershipService) Update(membership request.Membership) (request.Membership,error){
	if err := membership.Validate("update"); err!=nil{
		return membership ,err
	}
	m_entity := membership.ToDB()
	err:=service.repository.Update(&m_entity)
	res := request.DBtoReqMember(m_entity)
	return res ,err
}

func (service *membershipService) Delete(membership request.Membership) (request.Membership,error){
	if err := membership.Validate("delete"); err!=nil{
		return membership ,err
	}
	m_entity := membership.ToDB()
	err:=service.repository.Delete(&m_entity)
	res := request.DBtoReqMember(m_entity)
	return res ,err
}
