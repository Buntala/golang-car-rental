package repository

import (
	"car-rental/db"
	"car-rental/entity"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MembershipRepository interface {
	Save(membership entity.Membership) (entity.Membership,error)
	Update(membership entity.Membership) (entity.Membership,error)
	Delete(membership entity.Membership) (entity.Membership,error)
	FindAll() []entity.Membership
	FindOne(membership entity.Membership) (entity.Membership,error)
}

type databaseMembership struct {
	connection *gorm.DB
}

func NewMembershipRepository() MembershipRepository {
	return &databaseMembership{
		connection: db.DB,
	}
}

func (db *databaseMembership) Save(membership entity.Membership)(entity.Membership,error){
	status := db.connection.Clauses(clause.Returning{}).Create(&membership)
	if status.RowsAffected == 0{
		return membership, errors.New("no data with the id")
	}
	return membership,status.Error
}

func (db *databaseMembership) Update(membership entity.Membership) (entity.Membership,error){
	status := db.connection.Updates(&membership)
	if status.RowsAffected == 0{
		return membership, errors.New("no data with the id")
	}
	return membership,status.Error
}

func (db *databaseMembership) Delete(membership entity.Membership) (entity.Membership,error){
	status := db.connection.Delete(&membership)
	if status.RowsAffected == 0{
		return membership, errors.New("no data with the id")
	}
	return membership,status.Error
}

func (db *databaseMembership) FindAll() []entity.Membership {
	var memberships []entity.Membership
	db.connection.Set("gorm:auto_preload", true).Order("membership_id desc").Find(&memberships)
	return memberships
}

func (db *databaseMembership) FindOne(membership entity.Membership) (entity.Membership,error){
	status := db.connection.Find(&membership,membership.MembershipID)
	if status.RowsAffected == 0{
		return membership, errors.New("no data with the id")
	}
	return membership, status.Error
}
/*
func (db *databaseMembership) GetID(membership entity.Membership) int{
	var result entity.Membership
	status := db.connection.Where("name = ?", membership.Name).First(&result)
	return result.MembershipID
}*/
