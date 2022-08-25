package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConnect() *sqlx.DB {
	godotenv.Load()
	var (
		host     = os.Getenv("PGHOST")
		port     = os.Getenv("PGPORT")
		user     = os.Getenv("PGuser")
		password = os.Getenv("PGPASSWORD")
		dbname   = os.Getenv("PGDATABASE")
	)
	credStr := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s", host, user, dbname, password, port)
	fmt.Println(credStr)
	db, err := sqlx.Connect("postgres", credStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func DbConnectGorm() *gorm.DB {
	godotenv.Load()
	var (
		host     =  os.Getenv("PGHOST")
		port     = os.Getenv("PGPORT")
		user     = os.Getenv("PGuser")
		password = os.Getenv("PGPASSWORD")
		dbname   =  os.Getenv("PGDATABASE")
	)
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s", host,user,dbname,password,port)
	db, err :=gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}