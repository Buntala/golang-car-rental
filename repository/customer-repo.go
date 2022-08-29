package repository

import (
	"car-rental/db"
	"car-rental/entity"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CustomerRepository interface {
	Save(customer entity.Customer) (entity.Customer,error)
	Update(customer entity.Customer) (entity.Customer,error)
	Delete(customer entity.Customer) (entity.Customer,error)
	FindAll() []entity.Customer
	FindOne(customer entity.Customer) (entity.Customer,error)
	SaveMembership(customer entity.Customer)  (entity.Customer,error)

	GetMembershipID(membership entity.Membership) int
	GetMembershipName(membership entity.Membership) string
}

type database struct {
	connection *gorm.DB
}

func NewCustomerRepository() CustomerRepository {
	return &database{
		connection: db.DB,
	}
}

func (db *database) Save(customer entity.Customer) (entity.Customer,error){
	if customer.MembershipID ==0{
		status := db.connection.Clauses(clause.Returning{}).Omit("membership_id").Create(&customer)
		return customer, status.Error
	}
	status := db.connection.Clauses(clause.Returning{}).Create(&customer)
	return customer, status.Error
}

func (db *database) Update(customer entity.Customer) (entity.Customer,error){
	if customer.MembershipID ==0{
		status := db.connection.Clauses(clause.Returning{}).Omit("membership_id").Updates(&customer)
		if status.RowsAffected == 0{
			return customer, errors.New("no data with the id")
		}
		return customer, status.Error
	}
	status := db.connection.Clauses(clause.Returning{}).Updates(&customer)
	if status.RowsAffected == 0{
		return customer, errors.New("no data with the id")
	}
	return customer, status.Error
}

func (db *database) Delete(customer entity.Customer) (entity.Customer,error){
	status := db.connection.Clauses(clause.Returning{}).Omit("membership_id").Delete(&customer)
	if status.RowsAffected == 0{
		return customer, errors.New("no data with the id")
	}
	return customer, status.Error
}

func (db *database) FindAll() []entity.Customer {
	var customers []entity.Customer
	db.connection.Set("gorm:auto_preload", true).Order("customer_id desc").Find(&customers)
	return customers
}

func (db *database) FindOne(customer entity.Customer) (entity.Customer,error) {
	var customerRes entity.Customer
	status := db.connection.First(&customerRes,customer.CustomerID)
	if status.RowsAffected == 0{
		return customer, errors.New("no data with the id")
	}
	return customerRes , status.Error
}
func (db *database) SaveMembership(customer entity.Customer) (entity.Customer,error){
	status := db.connection.Clauses(clause.Returning{}).Model(&customer).Update("membership_id",customer.MembershipID)
	if status.RowsAffected == 0{
		return customer, errors.New("no data with the id")
	}
	return customer , status.Error
}
/*]
func (db *database) SaveMembership(membership entity.Membership) int{
	var result entity.Membership
	db.connection.Where("name = ?", membership.Name).First(&result)
	return result.MembershipID
}*/

func (db *database) GetMembershipID(membership entity.Membership) int{
	var result entity.Membership
	db.connection.Where("name = ?", membership.Name).First(&result)
	return result.MembershipID
}

func (db *database) GetMembershipName(membership entity.Membership) string{
	var result entity.Membership
	db.connection.First(&result,membership.MembershipID)
	return result.Name
}
