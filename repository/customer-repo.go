package repository

import (
	"car-rental/entity"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	Save(Customer entity.Customer)
	Update(Customer entity.Customer)
	Delete(Customer entity.Customer)
	FindAll() []entity.Customer
}

type database struct {
	connection *gorm.DB
}

func NewCustomerRepository() CustomerRepository {
	godotenv.Load()
	var (
		host     = os.Getenv("PGHOST")
		port     = os.Getenv("PGPORT")
		user     = os.Getenv("PGuser")
		password = os.Getenv("PGPASSWORD")
		dbname   = os.Getenv("PGDATABASE")
	)
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s", host, user, dbname, password, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	//db.AutoMigrate(&entity.Customer{}, &entity.Person{})
	return &database{
		connection: db,
	}
}

func (db *database) Save(customer entity.Customer) {
	db.connection.Create(&customer)
}

func (db *database) Update(customer entity.Customer) {
	db.connection.Updates(&customer)
}

func (db *database) Delete(customer entity.Customer) {
	db.connection.Delete(&customer)
}

func (db *database) FindAll() []entity.Customer {
	var customers []entity.Customer
	db.connection.Set("gorm:auto_preload", true).Order("customer_id desc").Find(&customers)
	return customers
}

