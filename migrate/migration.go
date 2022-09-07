package main

import (
	"car-rental/entity"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load()
	var (
		host     = os.Getenv("PGHOST")
		port     = os.Getenv("PGPORT")
		user     = os.Getenv("PGUSER")
		password = os.Getenv("PGPASSWORD")
		dbname   = os.Getenv("PGDATABASE")
	)
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s", host, user, dbname, password, port)
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true,})
	if err!=nil{
		log.Panic(err)
	}
	erro := DB.AutoMigrate(&entity.BookingType{},&entity.Car{},&entity.Customer{},&entity.Driver{},&entity.Membership{})
	if erro!= nil{
		log.Panic(erro.Error())
	}
	log.Print("Migration is Succesful")
}